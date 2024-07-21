package gapi

import (
	"fmt"
	db "remind/user/db/sqlc"
	"remind/user/pb"
	"remind/user/token"
	"remind/user/util"
	"remind/user/worker"
)

type Server struct {
	pb.UnimplementedUserRemindServer
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
