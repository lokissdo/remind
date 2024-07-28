package gapi

import (
	"context"
	"encoding/base64"
	"fmt"
	db "remind/journal/db/sqlc"
	"remind/journal/pb"

	"github.com/jackc/pgx/v5/pgtype"
)

func (server *Server) UpdateJournal(ctx context.Context, req *pb.UpdateJournalRequest) (*pb.UpdateJournalResponse, error) {
	if err := server.authorization(ctx, req.GetUsername()); err != nil {
		return nil, fmt.Errorf("authorization failed: %v", err)
	}

	arg := db.UpdateJournalParams{
		ID:       req.GetId(),
		Content: pgtype.Text{
			String: req.GetContent(),
			Valid:  req.Content != nil,
		},
		Status: pgtype.Bool{
			Bool:  req.GetStatus(),
			Valid: req.Status != nil,
		},
		Title: pgtype.Text{
			String: req.GetTitle(),
			Valid:  req.Title != nil,
		},
	}

	journal, err := server.store.UpdateJournal(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("failed to update journal: %v", err)
	}

	return &pb.UpdateJournalResponse{
		Journal: convertJournal(journal),
	}, nil
}

func (server *Server) AddImage(ctx context.Context, req *pb.AddImageRequest) (*pb.AddImageResponse, error) {
	if err := server.authorization(ctx, req.GetUsername()); err != nil {
		return nil, fmt.Errorf("authorization failed: %v", err)
	}

	imageBytes, err := base64.StdEncoding.DecodeString(req.GetContent())
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	arg := db.CreateImageParams{
		JournalID: req.GetJournalId(),
		Content:   imageBytes,
	}

	image, err := server.store.CreateImage(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("failed to create image: %v", err)
	}

	return &pb.AddImageResponse{
		Image: &pb.Image{
			Id:        image.ID,
			JournalId: image.JournalID,
			Content:   req.GetContent(),
		},
	}, nil
}

func (server *Server) AddAudio(ctx context.Context, req *pb.AddAudioRequest) (*pb.AddAudioResponse, error) {
	if err := server.authorization(ctx, req.GetUsername()); err != nil {
		return nil, fmt.Errorf("authorization failed: %v", err)
	}

	audioBytes, err := base64.StdEncoding.DecodeString(req.GetContent())
	if err != nil {
		return nil, fmt.Errorf("failed to decode audio: %v", err)
	}

	arg := db.CreateAudioParams{
		JournalID: req.GetJournalId(),
		Content:   audioBytes,
	}

	audio, err := server.store.CreateAudio(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("failed to create audio: %v", err)
	}

	return &pb.AddAudioResponse{
		Audio: &pb.Audio{
			Id:        audio.ID,
			JournalId: audio.JournalID,
			Content:  req.GetContent(),
		},
	}, nil
}