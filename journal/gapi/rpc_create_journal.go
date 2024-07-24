package gapi

import (
	"context"
	db "remind/journal/db/sqlc"
	"remind/journal/pb"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateJournal(ctx context.Context, req *pb.CreateJournalRequest) (*pb.CreateJournalResponse, error) {
	if err := server.authorization(ctx, req.GetUsername()); err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "authorization failed: %v", err)
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
			if err != nil {
				continue
			}
			arg := db.CreateImageParams{
				JournalID: journal.ID,
				Content: image,
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
			if err != nil {
				continue
			}
			arg := db.CreateAudioParams{
				JournalID: journal.ID,
				Content: audio,
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
