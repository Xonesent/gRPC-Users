package server

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"users/config"
)

type Server struct {
	cfg  *config.Config
	pgDB *sqlx.DB
	grpc *grpc.Server
}

func NewServer(
	cfg *config.Config,
	pgDB *sqlx.DB,
) *Server {
	return &Server{
		cfg:  cfg,
		pgDB: pgDB,
		grpc: grpc.NewServer(),
	}
}

func (s *Server) Run() error {
	err := s.MapHandlers()
	if err != nil {
		return err
	}

	go func() {
		listener, err := net.Listen("tcp", s.cfg.GRPCServer.Host)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		log.Println("Grpc Server is started ", s.cfg.GRPCServer.Host)
		defer listener.Close()

		if err := s.grpc.Serve(listener); err != nil {
			log.Fatalf("Failed to GRPC serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-quit

	s.grpc.GracefulStop()

	return nil
}

func NewDB(cfg config.ConfigPg) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
