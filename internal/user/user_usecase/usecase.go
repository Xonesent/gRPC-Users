package user_usecase

import (
	"context"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"users/internal/user"
	"users/pkg/models"
)

type UserUC struct {
	userPGRepo user.Repository
	trManager  *manager.Manager
}

func NewUserUC(clientPGR user.Repository, trManager *manager.Manager) *UserUC {
	return &UserUC{userPGRepo: clientPGR, trManager: trManager}
}

func (t *UserUC) AddUser(ctx context.Context, UserParams models.GrpcAddUser) (int64, error) {
	var UserId int64

	UserId, err := t.userPGRepo.CreateUser(ctx, UserParams)
	if err != nil {
		return -1, err
	}

	return UserId, nil
}

func (t *UserUC) GetUser(ctx context.Context, UserParams models.GrpcGetUser) (string, error) {
	userStr, err := t.userPGRepo.GetUser(ctx, UserParams)

	if err != nil {
		return "", err
	}

	return userStr, nil
}
