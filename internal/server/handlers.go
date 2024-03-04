package server

import (
	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"users/internal/user/user_delivery/grpc"
	"users/internal/user/user_repository"
	"users/internal/user/user_usecase"
	"users/pkg/api/userProto"
)

func (s *Server) MapHandlers() (err error) {
	UserPGR := user_repository.NewClientPGRepository(s.pgDB, trmsqlx.DefaultCtxGetter)

	trManager := manager.Must(trmsqlx.NewDefaultFactory(s.pgDB))
	UserUc := user_usecase.NewUserUC(UserPGR, trManager)

	userProtoServer := grpc.NewImplementation(UserUc)
	userProto.RegisterUserServiceServer(s.grpc, userProtoServer)

	return nil
}
