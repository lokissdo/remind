package rabbitmq

import (
	"errors"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

const (
	_retryTimes     = 5
	_backOffSeconds = 2
)

type RabbitMQConnStr string

var ErrCannotConnectRabbitMQ = errors.New("cannot connect to rabbit")

func NewRabbitMQConn(rabbitMqURL RabbitMQConnStr) (*amqp.Connection, error) {
	var (
		amqpConn *amqp.Connection
		counts   int64
	)

	for {
		connection, err := amqp.Dial(string(rabbitMqURL))
		if err != nil {
			log.Fatal().Err(err).Str("url", string(rabbitMqURL)).Msg("failed to connect to RabbitMq...")
			counts++
		} else {
			amqpConn = connection

			break
		}

		if counts > _retryTimes {
			log.Fatal().Err(err).Msg("failed to retry")

			return nil, ErrCannotConnectRabbitMQ
		}

		log.Info().Str("url", string(rabbitMqURL)).Msg("retrying to connect to RabbitMq...")
		time.Sleep(_backOffSeconds * time.Second)

		continue
	}

	log.Info().Str("url", string(rabbitMqURL)).Msg("connected to RabbitMq...")

	return amqpConn, nil
}