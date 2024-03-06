package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"users/internal/user"
	"users/pkg/api/userProto"
	"users/pkg/models"
)

type Implementation struct {
	userProto.UnimplementedUserServiceServer
	userUC user.UseCase
}

func NewImplementation(clientUC user.UseCase) *Implementation {
	return &Implementation{userUC: clientUC}
}

func (t *Implementation) AddUser(ctx context.Context, grpcParams *userProto.PostUserRequest,
) (*userProto.PostUserResponse, error) {
	UserParams := models.GrpcAddUser{
		Person: grpcParams.Person,
	}

	addedUserId, err := t.userUC.AddUser(ctx, UserParams)
	if err != nil {
		log.Print(err.Error())
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &userProto.PostUserResponse{
		Message: "Successfully added",
		UserId:  addedUserId,
	}, nil
}

func (t *Implementation) GetUser(ctx context.Context, grpcParams *userProto.GetUserRequest) (*userProto.GetUserResponse, error) {
	UserParams := models.GrpcGetUser{
		UserId: grpcParams.UserId,
	}

	gotUser, err := t.userUC.GetUser(ctx, UserParams)
	if err != nil {
		log.Print(err.Error())
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &userProto.GetUserResponse{
		Message: "Successfully got",
		Person:  gotUser,
	}, nil
}
