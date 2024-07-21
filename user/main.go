package main

import (
	"context"
	"errors"
	db "remind/user/db/sqlc"
	"remind/user/gapi"
	"remind/user/mail"
	"remind/user/pb"
	"remind/user/util"
	"remind/user/worker"

	// "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// "google.golang.org/protobuf/encoding/protojson"
	"net"
	// "net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "dev" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(connPool)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
	go runTaskProcessor(config, redisOpt, store)
	// go runGatewayServer(config, store, taskDistributor)
	runGrpcServer(config, store, taskDistributor)
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create migration")
	}

	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

func runTaskProcessor(config util.Config, redisOpt asynq.RedisClientOpt, store db.Store) {
	mailer := mail.NewGmailSender(config.EmailConfig.EmailSenderName, config.EmailConfig.EmailSenderAddress, config.EmailConfig.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Info().Msg("starting task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}

func runGrpcServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterUserRemindServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.ServerAddr.GRPC)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}

// func runGatewayServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) {
// 	server, err := gapi.NewServer(config, store, taskDistributor)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("cannot create server")
// 	}

// 	jsonOptions := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
// 		MarshalOptions: protojson.MarshalOptions{
// 			UseProtoNames: true,
// 		},
// 		UnmarshalOptions: protojson.UnmarshalOptions{
// 			DiscardUnknown: true,
// 		},
// 	})

// 	grpcMux := runtime.NewServeMux(jsonOptions)
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	err = pb.RegisterAuthRemindHandlerServer(ctx, grpcMux, server)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("cannot register gateway server")
// 	}

// 	mux := http.NewServeMux()
// 	mux.Handle("/", grpcMux)

// 	listener, err := net.Listen("tcp", config.ServerAddr.HTTP)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("cannot create listener")
// 	}

// 	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())
// 	handler := gapi.HttpLogger(mux)
// 	err = http.Serve(listener, handler)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("cannot start HTTP gateway server")
// 	}
// }
