package user

import (
	"context"
	"users/pkg/models"
)

type Repository interface {
	CreateUser(ctx context.Context, UserParams models.GrpcAddUser) (int64, error)
	GetUser(ctx context.Context, UserParams models.GrpcGetUser) (string, error)
}
