package gapi

import (
	"context"
	"encoding/base64"
	"encoding/json"
	db "remind/journal/db/sqlc"
	events "remind/journal/event"
	"remind/journal/pb"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateJournal(ctx context.Context, req *pb.CreateJournalRequest) (*pb.CreateJournalResponse, error) {
	// if err := server.authorization(ctx, req.GetUsername()); err != nil {
	// 	return nil, status.Errorf(codes.PermissionDenied, "authorization failed: %v", err)
	// }

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
	event := events.JournalCreated{
		Type:    "journal",
		ID:      journal.ID,
		Content: journal.Content.String,
		Username: req.GetUsername(),
	}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		return nil, errors.Wrap(err, "json.Marshal[event]")
	}
	server.Publisher.Publish(ctx, eventBytes, "text/plain")

	imageList := req.GetImages()
	insertedImages := make([]*pb.Image, 0, len(imageList))
	if len(imageList) > 0 {
		for _, image := range imageList {
			imageBytes, err := base64.StdEncoding.DecodeString(image)
			if err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "failed to decode image: %v", err)
			}
			arg := db.CreateImageParams{
				JournalID: journal.ID,
				Content:   imageBytes,
			}
			dbImage, err := server.store.CreateImage(ctx, arg)
			if err == nil {
				insertedImage := &pb.Image{
					JournalId: journal.ID,
					Content:   image,
					CreatedAt: timestamppb.New(journal.CreatedAt),
				}
				insertedImages = append(insertedImages, insertedImage)

				imageEvent := events.ImageCreated{
					Type:      "image",
					ID:        dbImage.ID,
					JournalID: dbImage.JournalID,
					Content:   base64.StdEncoding.EncodeToString(dbImage.Content),
					Username:  req.GetUsername(),
				}
				imageEventBytes, err := json.Marshal(imageEvent)
				if err != nil {
					return nil, errors.Wrap(err, "json.Marshal[imageEvent]")
				}
				server.Publisher.Publish(ctx, imageEventBytes, "text/plain")
			}
		}
	}

	audioList := req.GetAudios()
	insertedAudios := make([]*pb.Audio, 0, len(audioList))
	if len(audioList) > 0 {
		for _, audio := range audioList {
			audioBytes, err := base64.StdEncoding.DecodeString(audio)
			if err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "failed to decode audio: %v", err)
			}
			arg := db.CreateAudioParams{
				JournalID: journal.ID,
				Content:   audioBytes,
			}
			_, err = server.store.CreateAudio(ctx, arg)
			if err == nil {
				insertedAudio := &pb.Audio{
					JournalId: journal.ID,
					Content:   audio,
					CreatedAt: timestamppb.New(journal.CreatedAt),
				}
				insertedAudios = append(insertedAudios, insertedAudio)
			}
		}
	}

	return &pb.CreateJournalResponse{
		Journal: convertJournal(journal),
		Images:  insertedImages,
		Audios:  insertedAudios,
	}, nil
}
