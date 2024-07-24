package gapi

import (
	"context"
	"fmt"
	pb "remind/journal/pb"
)

func (server *Server) DeleteJournal(ctx context.Context, req *pb.DeleteJournalRequest) (*pb.DeleteJournalResponse, error) {
	if err := server.authorization(ctx, req.GetUsername()); err != nil {
		return nil, fmt.Errorf("authorization failed: %v", err)
	}

	err := server.store.DeleteJournalTx(ctx, req.GetId())
	if err != nil {
		return nil, fmt.Errorf("cannot delete journal: %v", err)
	}
	
	return &pb.DeleteJournalResponse{
		Success: true,
	}, nil
}

func (server *Server) DeleteImage(ctx context.Context, req *pb.DeleteImageRequest) (*pb.DeleteImageResponse, error) {
	if err := server.authorization(ctx, req.GetUsername()); err != nil {
		return nil, fmt.Errorf("authorization failed: %v", err)
	}

	err := server.store.DeleteImage(ctx, req.GetId())
	if err != nil {
		return nil, fmt.Errorf("cannot delete image: %v", err)
	}
	
	return &pb.DeleteImageResponse{
		Success: true,
	}, nil
}

func (server *Server) DeleteAudio(ctx context.Context, req *pb.DeleteAudioRequest) (*pb.DeleteAudioResponse, error) {
	if err := server.authorization(ctx, req.GetUsername()); err != nil {
		return nil, fmt.Errorf("authorization failed: %v", err)
	}

	err := server.store.DeleteAudio(ctx, req.GetId())
	if err != nil {
		return nil, fmt.Errorf("cannot delete audio: %v", err)
	}
	
	return &pb.DeleteAudioResponse{
		Success: true,
	}, nil
}