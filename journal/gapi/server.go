package gapi

import (
	"context"
	"fmt"
	"remind/journal/db/sqlc"
	"remind/journal/pb"
	"remind/journal/util"

	"google.golang.org/grpc/metadata"
)

type Server struct {
	pb.UnimplementedJournalRemindServer
	config          util.Config
	store           db.Store
	authorization   func(ctx context.Context, username string) error 
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
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
	}
	return server, nil
}
