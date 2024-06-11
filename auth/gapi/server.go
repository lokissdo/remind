package gapi

import (
	"fmt"
	db "github.com/elliotnguyen/auth/db/sqlc"
	"github.com/elliotnguyen/auth/pb"
	"github.com/elliotnguyen/auth/token"
	"github.com/elliotnguyen/auth/util"
	"github.com/elliotnguyen/auth/worker"
)

type Server struct {
	pb.UnimplementedAuthRemindServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.PasetoConfig.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	var server = &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}
	return server, nil
}
