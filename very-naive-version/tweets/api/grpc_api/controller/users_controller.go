package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"tweets/api/grpc_api/converter/follow_converter"
	"tweets/api/grpc_api/converter/user_converter"
	usersService "tweets/api/grpc_api/protobuf/users_service"
	"tweets/errs"
	"tweets/usecase"
)

type UsersController struct {
	usersService.UnimplementedUsersServiceServer
	usersUseCase usecase.UsersUseCase
}

func NewUsersController(usersUseCase usecase.UsersUseCase) UsersController {
	return UsersController{usersUseCase: usersUseCase}
}

func (controller *UsersController) SignUp(
	ctx context.Context, request *usersService.SignUpRequest,
) (*usersService.SignUpResponse, error) {

	userToBeCreated := user_converter.FromProtobufModelToDatabaseModel(
		user_converter.FromSignUpRequestToProtobufModel(request),
	)

	createdUser, err := controller.usersUseCase.CreateUser(userToBeCreated)
	if err != nil {
		if errors.As(err, &errs.ErrorAlreadyExists{}) {
			errAlreadyExists := err.(errs.ErrorAlreadyExists)
			return nil, status.Errorf(
				codes.AlreadyExists,
				"user already exists: unique fields with repeated values: %s",
				errAlreadyExists.FieldsWithRepeatedValues,
			)
		}

		slog.Error(
			"internal error when creating user",
			"error", err,
			"request", request.String(),
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	response := usersService.SignUpResponse{
		User: user_converter.FromDatabaseModelToProtobufModel(createdUser),
	}
	return &response, nil
}

func (controller *UsersController) FollowUser(
	ctx context.Context,
	request *usersService.FollowUserRequest,
) (*usersService.FollowUserResponse, error) {

	followerId, err := uuid.Parse(request.GetFollowerId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("invalid follower_id uuid: %s", err),
		)
	}

	followeeId, err := uuid.Parse(request.GetFolloweeId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("invalid followee_id uuid: %s", err),
		)
	}

	createdFollow, err := controller.usersUseCase.FollowUser(followerId, followeeId)
	if err != nil {
		if errors.As(err, &errs.ErrorNotFound{}) {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		if errors.As(err, &errs.ErrorUserAlreadyFollow{}) {
			return nil, status.Errorf(
				codes.AlreadyExists,
				"user with id %s already follow user with id %s",
				followerId, followeeId,
			)
		}

		slog.Error(
			"internal error when following user",
			"error", err,
			"request", request.String(),
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	response := usersService.FollowUserResponse{
		Follow: follow_converter.FromDatabaseModelToProtobufModel(createdFollow),
	}
	return &response, nil
}

func (controller *UsersController) UnfollowUser(
	ctx context.Context,
	request *usersService.UnfollowUserRequest,
) (*usersService.UnfollowUserResponse, error) {

	followerId, err := uuid.Parse(request.GetFollowerId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("invalid follower_id uuid: %s", err),
		)
	}

	followeeId, err := uuid.Parse(request.GetFolloweeId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("invalid followee_id uuid: %s", err),
		)
	}

	if err = controller.usersUseCase.UnfollowUser(followerId, followeeId); err != nil {
		if errors.As(err, &errs.ErrorNotFound{}) {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		if errors.As(err, &errs.ErrorUserDoesNotFollow{}) {
			return nil, status.Errorf(
				codes.NotFound,
				"user with id %s does not follow user with id %s",
				followerId, followeeId,
			)
		}

		slog.Error(
			"internal error when unfollowing",
			"error", err,
			"request", request.String(),
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	response := usersService.UnfollowUserResponse{}
	return &response, nil
}

func (controller *UsersController) DeleteUserById(
	ctx context.Context,
	request *usersService.DeleteUserByIdRequest,
) (*usersService.DeleteUserByIdResponse, error) {

	userId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id uuid: %s", err)
	}

	if err := controller.usersUseCase.DeleteUserById(userId); err != nil {
		if errors.As(err, &errs.ErrorNotFound{}) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		slog.Error(
			"internal error when deleting user by id",
			"error", err,
			"request", request.String(),
		)
		return nil, status.Error(codes.Internal, "internal error")
	}

	response := usersService.DeleteUserByIdResponse{}
	return &response, nil
}

func (controller *UsersController) GetUserById(
	ctx context.Context,
	request *usersService.GetUserByIdRequest,
) (*usersService.GetUserByIdResponse, error) {

	userId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("invalid user_id uuid: %s", err),
		)
	}

	user, err := controller.usersUseCase.GetUserById(userId)
	if err != nil {
		if errors.As(err, &errs.ErrorNotFound{}) {
			return nil, status.Errorf(
				codes.NotFound, err.Error(),
			)
		}

		slog.Error(
			"internal error when getting user by id",
			"error", err,
			"request", request.String(),
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	response := usersService.GetUserByIdResponse{
		User: user_converter.FromDatabaseModelToProtobufModel(user),
	}
	return &response, nil
}
func (controller *UsersController) GetAllUsers(
	ctx context.Context,
	request *usersService.GetAllUsersRequest,
) (*usersService.GetAllUsersResponse, error) {

	allUsers, err := controller.usersUseCase.GetAllUsers()
	if err != nil {
		slog.Error(
			"internal error when getting all users",
			"error", err,
			"request", request.String(),
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	response := usersService.GetAllUsersResponse{
		Users: user_converter.FromDatabaseModelsToProtobufModels(allUsers),
	}
	return &response, nil
}
