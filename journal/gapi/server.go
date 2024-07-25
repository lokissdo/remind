package gapi

import (
	"context"
	// "encoding/json"
	"fmt"
	"log/slog"
	"remind/journal/db/sqlc"
	"remind/journal/pb"
	"remind/journal/util"

	// event "remind/journal/event"
	pkgPublisher "remind/pkg/rabbitmq/publisher"
	pkgConsumer "remind/pkg/rabbitmq/consumer"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	pb.UnimplementedJournalRemindServer
	config          util.Config
	store           db.Store
	authorization   func(ctx context.Context, username string) error
	AMQPConn  		*amqp.Connection
	Publisher 		pkgPublisher.EventPublisher
	Consumer 		pkgConsumer.EventConsumer
}

func NewServer(
	config util.Config, store db.Store, amqpConn *amqp.Connection, 
	publisher pkgPublisher.EventPublisher, consumer pkgConsumer.EventConsumer,
) (*Server, error) {
	authorization := func (ctx context.Context, username string) error {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return fmt.Errorf("metadata is not provided")
		}
		values := md.Get("username")
		if len(values) == 0 {
			return fmt.Errorf("username is not provided")
		}
		if values[0] != username {
			return fmt.Errorf("user %s is not allowed to create journal for user %s", values[0], username)
		}
		return nil
	}

	var server = &Server{
		config:          config,
		store:           store,
		authorization:   authorization,
		AMQPConn:        amqpConn,
		Publisher:       publisher,
		Consumer:        consumer,
	}

	return server, nil
}

func (s *Server) Worker(ctx context.Context, messages <-chan amqp.Delivery) {
	for delivery := range messages {
		slog.Info("processDeliveries", "delivery_tag", delivery.DeliveryTag)
		slog.Info("received", "delivery_type", delivery.Type)

		// switch delivery.Type {
		// case "image-embedded":
		// 	var payload event.ImageEmbedded

		// 	err := json.Unmarshal(delivery.Body, &payload)
		// 	if err != nil {
		// 		slog.Error("failed to Unmarshal message", err)
		// 	}

		// 	arg := db.UpdateImageEmbeddingStatusParams{
		// 		ID:       	payload.ID,
		// 		IsEmbedded: true,
		// 	}

		// 	_, err = s.store.UpdateImageEmbeddingStatus(ctx, arg)
		// 	if err != nil {
		// 		if err = delivery.Reject(false); err != nil {
		// 			slog.Error("failed to delivery.Reject", err)
		// 		}

		// 		slog.Error("failed to process delivery", err)
		// 	} else {
		// 		err = delivery.Ack(false)
		// 		if err != nil {
		// 			slog.Error("failed to acknowledge delivery", err)
		// 		}
		// 	}
		// case "journal-updated":
		// 	var payload event.JournalEmbedded

		// 	err := json.Unmarshal(delivery.Body, &payload)
		// 	if err != nil {
		// 		slog.Error("failed to Unmarshal message", err)
		// 	}

		// 	arg := db.UpdateJournalParams{
		// 		ID:        	payload.ID,
		// 		IsEmbedded: true,
		// 	}

		// 	_, err = s.store.UpdateJournal(ctx, arg)

		// 	if err != nil {
		// 		if err = delivery.Reject(false); err != nil {
		// 			slog.Error("failed to delivery.Reject", err)
		// 		}

		// 		slog.Error("failed to process delivery", err)
		// 	} else {
		// 		err = delivery.Ack(false)
		// 		if err != nil {
		// 			slog.Error("failed to acknowledge delivery", err)
		// 		}
		// 	}
		// default:
		// 	slog.Info("default")
		// }
	}

	slog.Info("Deliveries channel closed")
}