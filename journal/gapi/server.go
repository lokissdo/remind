package gapi

import (
	db "remind/journal/db/sqlc"
	"remind/journal/pb"
	"remind/journal/util"
)

type Server struct {
	pb.UnimplementedJournalRemindServer
	config          util.Config
	store           db.Store
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	var server = &Server{
		config:          config,
		store:           store,
	}
	return server, nil
}
