package main

import (
	"net"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"remind/auth/gapi"
	"remind/auth/pb"
	"remind/auth/util"

	"remind/pkg/logger"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "dev" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	runGrpcServer(config)
}

func runGrpcServer(config util.Config) {
	server, err := gapi.NewServer(config)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	grpcLogger := grpc.UnaryInterceptor(logger.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterAuthRemindServer(grpcServer, server)
	reflection.Register(grpcServer)

	listen, err := net.Listen("tcp", config.GRPC)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}

	log.Info().Str("address", config.GRPC).Msg("server is running")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}