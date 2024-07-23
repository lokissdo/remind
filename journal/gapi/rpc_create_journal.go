package gapi

import (
	"context"
	"encoding/base64"
	"fmt"
	_ "fmt"
	db "remind/journal/db/sqlc"
	"remind/journal/pb"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateJournal(ctx context.Context, req *pb.CreateJournalRequest) (*pb.CreateJournalResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata is not provided")
	}

	log.Info().Msgf("metadata: %v", md)

	values := md.Get("username")
	if len(values) == 0 {
		return nil, fmt.Errorf("username is not provided")
	}
	username := values[0]
	if username != req.GetUsername() {
		return nil, fmt.Errorf("user %s is not allowed to create journal for user %s", username, req.GetUsername())
	}

	arg := db.CreateJournalParams{
		Username: req.GetUsername(),
		Title:    req.GetTitle(),
		Content: pgtype.Text{
			String: req.GetContent(),
			Valid:  true,
		},
		Status: req.GetStatus(),
	}

	journal, err := server.store.CreateJournal(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create journal: %v", err)
	}

	imageList := req.GetImages()
	insertedImages := make([]*pb.Image, 0, len(imageList))
	if len(imageList) > 0 {
		for _, image := range imageList {
			imageBytes, err := base64.StdEncoding.DecodeString(image)
			if err != nil {
				continue
			}
			arg := db.CreateImageParams{
				JournalID: journal.ID,
				Content: imageBytes,
			}
			_, err = server.store.CreateImage(ctx, arg)
			if err == nil {
				insertedImage := &pb.Image{
					JournalId: journal.ID,
					Content: image,
					CreatedAt: timestamppb.New(journal.CreatedAt),
				}
				insertedImages = append(insertedImages, insertedImage)
			}
		}
	}

	audioList := req.GetAudios()
	insertedAudios := make([]*pb.Audio, 0, len(audioList))
	if len(audioList) > 0 {
		for _, audio := range audioList {
			audioBytes, err := base64.StdEncoding.DecodeString(audio)
			if err != nil {
				continue
			}
			arg := db.CreateAudioParams{
				JournalID: journal.ID,
				Content: audioBytes,
			}
			_, err = server.store.CreateAudio(ctx, arg)
			if err == nil {
				insertedAudio := &pb.Audio{
					JournalId: journal.ID,
					Content: audio,
					CreatedAt: timestamppb.New(journal.CreatedAt),
				}
				insertedAudios = append(insertedAudios, insertedAudio)
			}
		}
	}

	return &pb.CreateJournalResponse{
		Journal: convertJournal(journal),
		Images: insertedImages,
		Audios: insertedAudios,
	}, nil
}

func convertJournal(journal db.Journal) (*pb.Journal) {
	return &pb.Journal{
		Id:       journal.ID,
		Username: journal.Username,
		Title:    journal.Title,
		Content:  journal.Content.String,
		Status:   journal.Status,
		CreatedAt: timestamppb.New(journal.CreatedAt),
		UpdatedAt: timestamppb.New(journal.UpdatedAt),
	}
}
