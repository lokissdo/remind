package gapi

import (
	"encoding/base64"
	db "remind/journal/db/sqlc"
	"remind/journal/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

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

func convertImage(image db.Image) *pb.Image {
	encodedImage := base64.StdEncoding.EncodeToString(image.Content)
	return &pb.Image{
		Id:        image.ID,
		JournalId: image.JournalID,
		Content:   encodedImage,
		CreatedAt: timestamppb.New(image.CreatedAt),
	}
}

func convertAudio(audio db.Audio) *pb.Audio {
	encodedAudio := base64.StdEncoding.EncodeToString(audio.Content)
	return &pb.Audio{
		Id:        audio.ID,
		JournalId: audio.JournalID,
		Content:   encodedAudio,
		CreatedAt: timestamppb.New(audio.CreatedAt),
	}
}