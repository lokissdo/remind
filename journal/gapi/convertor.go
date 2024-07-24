package gapi

import (
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
	return &pb.Image{
		Id:        image.ID,
		JournalId: image.JournalID,
		Content:   image.Content,
		CreatedAt: timestamppb.New(image.CreatedAt),
	}
}

func convertAudio(audio db.Audio) *pb.Audio {
	return &pb.Audio{
		Id:        audio.ID,
		JournalId: audio.JournalID,
		Content:   audio.Content,
		CreatedAt: timestamppb.New(audio.CreatedAt),
	}
}