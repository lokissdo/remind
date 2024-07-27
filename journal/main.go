package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	db "remind/journal/db/sqlc"
	"remind/journal/gapi"

	"remind/journal/pb"
	"remind/journal/util"

	"remind/pkg/logger"
	"remind/pkg/rabbitmq"
	"remind/pkg/rabbitmq/consumer"
	"remind/pkg/rabbitmq/publisher"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	amqp, journalPublisher, journalConsumer, err := configMessageQueue(config.RabbitMQURL)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot configure message queue")
	}

	store := db.NewStore(connPool)

	runGrpcServer(config, store, amqp, journalPublisher, journalConsumer)
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

func configMessageQueue(rabbitmqURL string) (*amqp091.Connection, publisher.EventPublisher, consumer.EventConsumer, error) {
	amqp, err := rabbitmq.NewRabbitMQConn(rabbitmq.RabbitMQConnStr(rabbitmqURL))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create RabbitMQ connection")
	}

	journalPublisher, err := publisher.NewPublisher(amqp)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create RabbitMQ publisher")
	}
	journalPublisher.Configure(
		publisher.ExchangeName("journal-created-exchange"),
		publisher.BindingKey("journal-created-routing-key"),
		publisher.MessageTypeName("journal-created"),
	)

	journalConsumer, err := consumer.NewConsumer(amqp)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create RabbitMQ consumer")
	}
	journalConsumer.Configure(
		consumer.ExchangeName("journal-embedded-exchange"),
		consumer.BindingKey("journal-embedded-routing-key"),
		consumer.QueueName("journal-embedded"),
		consumer.ConsumerTag("journal-embedded-consumer"),
	)

	return amqp, journalPublisher, journalConsumer, nil
}

func runGrpcServer(config util.Config, store db.Store, amqp *amqp091.Connection, publisher publisher.EventPublisher, consumer consumer.EventConsumer) {
	ctx, cancel := context.WithCancel(context.Background())

	server, err := gapi.NewServer(config, store, amqp, publisher, consumer)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	grpcLogger := grpc.UnaryInterceptor(logger.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger, grpc.MaxRecvMsgSize(1024 * 1024 * 1024))
	pb.RegisterJournalRemindServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.ServerAddr)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	go func() {
		err1 := server.Consumer.StartConsumer(server.Worker)
		if err1 != nil {
			log.Fatal().Err(err1).Msg("cannot start consumer")
			cancel()
			<-ctx.Done()
		}
	}()

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}