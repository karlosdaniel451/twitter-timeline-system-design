package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"tweets/domain/models"
	"tweets/errs"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	FollowUser(followeeId, followerId uuid.UUID) (*models.Follow, error)
	UnfollowUser(followeeId, followerId uuid.UUID) error
	DoesUserFollow(followerId, followeeId uuid.UUID) (bool, error)
	DeleteUserById(id uuid.UUID) error
	GetUserById(id uuid.UUID) (*models.User, error)
	GetFollowerIdsOfUser(id uuid.UUID) ([]uuid.UUID, error)
	GetFolloweeIdsOfUser(id uuid.UUID) ([]uuid.UUID, error)
	GetAllUsers() ([]*models.User, error)
}

type UserRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) *UserRepositoryDB {
	return &UserRepositoryDB{db: db}
}

func (repository UserRepositoryDB) CreateUser(user *models.User) (*models.User, error) {
	result := repository.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repository UserRepositoryDB) FollowUser(
	followerId uuid.UUID,
	followeeId uuid.UUID,
) (*models.Follow, error) {

	follower, err := repository.GetUserById(followerId)
	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return nil, errs.ErrorNotFound{
				Msg: fmt.Sprintf("there is no user with id %s", followerId.String()),
			}
		}
		return nil, err
	}

	followee, err := repository.GetUserById(followeeId)
	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return nil, errs.ErrorNotFound{
				Msg: fmt.Sprintf("there is no user with id %s", followeeId.String()),
			}
		}
		return nil, err
	}

	follow := models.Follow{
		FollowerId: followerId,
		FolloweeId: followeeId,
	}

	// In case the Follow record was to be created was soft-deleted, that is,
	// it should be restored.
	result := repository.db.Unscoped().Model(&models.Follow{}).
		Where(&follow).Update("deleted_at", nil)

	if result.Error != nil {
		return nil, err
	}

	// If the Follow record was not soft-deleted.
	if result.RowsAffected != 1 {
		if err = repository.db.Create(&follow).Error; err != nil {
			return nil, err
		}
	}

	follower.PublicMetrics.FollowingCount++
	followee.PublicMetrics.FollowersCount++

	//err = repository.db.Model(followee).Where("id = ?", follower.Id).
	//	Update("followers_count", follower.PublicMetrics.FollowersCount).Error
	err = repository.db.Save(followee).Error
	if err != nil {
		return nil, fmt.Errorf("error when updating followers_count: %s", err)
	}

	//err = repository.db.Model(follower).Where("id = ?", followee.Id).
	//	Update("following_count", followee.PublicMetrics.FollowingCount).Error
	err = repository.db.Save(follower).Error
	if err != nil {
		return nil, fmt.Errorf("error when updating following_count: %s", err)
	}

	return &follow, nil
}

func (repository UserRepositoryDB) UnfollowUser(
	followerId uuid.UUID,
	followeeId uuid.UUID,
) error {

	follower, err := repository.GetUserById(followerId)
	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return errs.ErrorNotFound{
				Msg: fmt.Sprintf("there is no user with id %s", followeeId.String()),
			}
		}
		return err
	}

	followee, err := repository.GetUserById(followeeId)
	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return errs.ErrorNotFound{
				Msg: fmt.Sprintf("there is no user with id %s", followeeId.String()),
			}
		}
		return err
	}

	followToBeDeleted := models.Follow{
		FollowerId: followerId,
		FolloweeId: followeeId,
	}

	if err = repository.db.Delete(&followToBeDeleted).Error; err != nil {
		return err
	}

	follower.PublicMetrics.FollowingCount--
	followee.PublicMetrics.FollowersCount--

	err = repository.db.Save(followee).Error
	if err != nil {
		return fmt.Errorf("error when updating followers_count: %s", err)
	}

	err = repository.db.Save(follower).Error
	if err != nil {
		return fmt.Errorf("error when updating following_count: %s", err)
	}
	return nil
}

func (repository UserRepositoryDB) DoesUserFollow(
	followerId,
	followeeId uuid.UUID,
) (bool, error) {

	if err := repository.db.First(&models.Follow{
		FollowerId: followerId,
		FolloweeId: followeeId,
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (repository UserRepositoryDB) DeleteUserById(id uuid.UUID) error {
	var user models.User

	result := repository.db.First(&user, id)
	if result.Error != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return &errs.ErrorNotFound{
				Msg: fmt.Sprintf("there is no user with id %s", id.String()),
			}
		}
		return result.Error
	}

	if err := result.Delete(&user).Error; err != nil {
		return result.Error
	}

	return nil
}

func (repository UserRepositoryDB) GetUserById(id uuid.UUID) (*models.User, error) {
	var user models.User

	//result := repository.db.Preload("User.Followers").
	//	Preload("User.Followees").Find(&user, id)
	result := repository.db.First(&user, id)
	if result.Error != nil {
		if errors.As(result.Error, &gorm.ErrRecordNotFound) {
			return nil, errs.ErrorNotFound{
				Msg: fmt.Sprintf("there is no user with id %s", id.String()),
			}
		}
		return nil, result.Error
	}

	followersIds, err := repository.GetFollowerIdsOfUser(id)
	if err != nil {
		return nil, fmt.Errorf("error when getting follower ids of user: %s", err)
	}

	followeesIds, err := repository.GetFolloweeIdsOfUser(id)
	if err != nil {
		return nil, fmt.Errorf("error when getting followee ids of user: %s", err)
	}

	user.FollowerIds = followersIds
	user.FolloweeIds = followeesIds

	return &user, nil
}

func (repository UserRepositoryDB) GetFollowerIdsOfUser(id uuid.UUID) ([]uuid.UUID, error) {
	followersIds := make([]uuid.UUID, 0)

	err := repository.db.Model(&models.User{}).
		Joins("JOIN follows ON follows.follower_id = users.id").
		Where("follows.followee_id = ?", id).
		Select("follows.follower_id").Find(&followersIds).Error
	if err != nil {
		return nil, err
	}

	return followersIds, nil
}
func (repository UserRepositoryDB) GetFolloweeIdsOfUser(id uuid.UUID) ([]uuid.UUID, error) {
	followeesIds := make([]uuid.UUID, 0)

	err := repository.db.Model(&models.User{}).
		Joins("JOIN follows ON follows.follower_id = users.id").
		Where("follows.follower_id = ?", id).
		Select("follows.followee_id").Find(&followeesIds).Error
	if err != nil {
		return nil, err
	}

	return followeesIds, nil
}

func (repository UserRepositoryDB) GetAllUsers() ([]*models.User, error) {
	allUsers := make([]*models.User, 0)

	if err := repository.db.Find(&allUsers).Error; err != nil {
		return nil, err
	}

	var followersIds, followeesIds []uuid.UUID
	var err error

	for _, user := range allUsers {
		followersIds, err = repository.GetFollowerIdsOfUser(*user.Id)
		if err != nil {
			return nil, fmt.Errorf("error when getting follower ids of user: %s", err)
		}

		followeesIds, err = repository.GetFolloweeIdsOfUser(*user.Id)
		if err != nil {
			return nil, fmt.Errorf("error when getting followee ids of user: %s", err)
		}

		user.FollowerIds = followersIds
		user.FolloweeIds = followeesIds
	}

	return allUsers, nil
}
