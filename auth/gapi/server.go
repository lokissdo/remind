package gapi

import (
	"fmt"
	"remind/auth/pb"
	"remind/auth/token"
	"remind/auth/util"
)

type Server struct {
	pb.UnimplementedAuthRemindServer
	tokenMaker token.Maker
	config 	   util.Config
}

func NewServer(config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	var server = &Server{
		tokenMaker: tokenMaker,
		config: config,
	}

	return server, nil
}