package usecase

import (
	"fmt"
	"tweets/domain/models"
	"tweets/errs"
	"tweets/repository"

	"github.com/google/uuid"
)

type UsersUseCase interface {
	CreateUser(user *models.User) (*models.User, error)
	FollowUser(followerId, followeeId uuid.UUID) (*models.Follow, error)
	UnfollowUser(followerId, followeeId uuid.UUID) error
	DoesUserFollow(followerId, followeeId uuid.UUID) (bool, error)
	DeleteUserById(id uuid.UUID) error
	GetUserById(id uuid.UUID) (*models.User, error)
	GetAllUsers() ([]*models.User, error)

	CheckUserUniqueFields(
		id *uuid.UUID,
		username, email *string,
	) (*models.UserUniqueFieldsChecking, error)
}

type UsersUseCaseImpl struct {
	usersRepository repository.UserRepository
}

func NewUsersUseCaseImpl(usersRepository repository.UserRepository) *UsersUseCaseImpl {
	return &UsersUseCaseImpl{usersRepository: usersRepository}
}

func (useCase UsersUseCaseImpl) CreateUser(user *models.User) (*models.User, error) {
	uniqueFieldsChecking, err := useCase.CheckUserUniqueFields(
		nil, user.Username, user.Email,
	)
	if err != nil {
		return nil, fmt.Errorf("error when checking if user already exists: %s", err)
	}

	if uniqueFieldsChecking.Email || uniqueFieldsChecking.Username {
		var repeatedUniqueFields []string

		if uniqueFieldsChecking.Username {
			repeatedUniqueFields = append(repeatedUniqueFields, "username")
		}
		if uniqueFieldsChecking.Email {
			repeatedUniqueFields = append(repeatedUniqueFields, "email")
		}

		return nil, errs.ErrorAlreadyExists{
			Msg:                      fmt.Sprintf("user already exists"),
			FieldsWithRepeatedValues: repeatedUniqueFields,
		}
	}

	return useCase.usersRepository.CreateUser(user)
}

func (useCase UsersUseCaseImpl) FollowUser(
	followerId uuid.UUID,
	followeeId uuid.UUID,
) (*models.Follow, error) {

	followExist, err := useCase.DoesUserFollow(followerId, followeeId)
	if err != nil {
		return nil, fmt.Errorf("error when checking if follow already exists: %s", err)
	}

	if followExist {
		return nil, errs.ErrorUserAlreadyFollow{}
	}

	return useCase.usersRepository.FollowUser(followerId, followeeId)
}

func (useCase UsersUseCaseImpl) UnfollowUser(followerId, followeeId uuid.UUID) error {
	followExist, err := useCase.DoesUserFollow(followerId, followeeId)
	if err != nil {
		return fmt.Errorf("error when checking if follow already exists")
	}

	if !followExist {
		return errs.ErrorUserDoesNotFollow{}
	}

	return useCase.usersRepository.UnfollowUser(followerId, followeeId)
}

func (useCase UsersUseCaseImpl) DoesUserFollow(
	followerId,
	followeeId uuid.UUID,
) (bool, error) {

	return useCase.usersRepository.DoesUserFollow(followerId, followeeId)
}

func (useCase UsersUseCaseImpl) DeleteUserById(id uuid.UUID) error {
	return useCase.usersRepository.DeleteUserById(id)
}

func (useCase UsersUseCaseImpl) GetUserById(id uuid.UUID) (*models.User, error) {
	return useCase.usersRepository.GetUserById(id)
}

func (useCase UsersUseCaseImpl) GetAllUsers() ([]*models.User, error) {
	return useCase.usersRepository.GetAllUsers()
}

func (useCase UsersUseCaseImpl) CheckUserUniqueFields(
	id *uuid.UUID, username *string, email *string,
) (*models.UserUniqueFieldsChecking, error) {

	return useCase.usersRepository.CheckUserUniqueFields(id, username, email)
}
