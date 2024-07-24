package gapi

import (
	"context"
	"fmt"
	db "remind/journal/db/sqlc"
	pb "remind/journal/pb"
)

func (server *Server) QueryJournal(ctx context.Context, req *pb.QueryJournalRequest) (*pb.QueryJournalResponse, error) {
	if err := server.authorization(ctx, req.GetUsername()); err != nil {
		return nil, fmt.Errorf("authorization failed: %v", err)
	}

	journal, err := server.store.GetJournal(ctx, req.GetId())
	if err != nil {
		return nil, fmt.Errorf("cannot get journal: %v", err)
	}

	iamgeList, err := server.store.GetImageOfJournal(ctx, req.GetId())
	if err != nil {
		return nil, fmt.Errorf("cannot get image list: %v", err)
	}
	pbImageList := make([]*pb.Image, 0, len(iamgeList))
	for _, image := range iamgeList {
		pbImageList = append(pbImageList, convertImage(image))
	}

	audioList, err := server.store.GetAudioOfJournal(ctx, req.GetId())
	if err != nil {
		return nil, fmt.Errorf("cannot get audio list: %v", err)
	}
	pbAudioList := make([]*pb.Audio, 0, len(audioList))
	for _, audio := range audioList {
		pbAudioList = append(pbAudioList, convertAudio(audio))
	}

	return &pb.QueryJournalResponse{
		Journal: convertJournal(journal),
		Images: pbImageList,
		Audios: pbAudioList,
	}, nil
}

func (server *Server) QueryJournals(ctx context.Context, req *pb.QueryJournalsRequest) (*pb.QueryJournalsResponse, error) {
	if err := server.authorization(ctx, req.GetUsername()); err != nil {
		return nil, fmt.Errorf("authorization failed: %v", err)
	}

	limit := req.GetLimit()
	if limit == 0 {
		limit = 10
	}

	arg := db.QueryJournalParams{
		Username: req.GetUsername(),
		ToTsquery: req.GetSearch(),
		Limit:   limit,
		Offset:  req.GetOffset(),
	}

	journalList, err := server.store.QueryJournal(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("cannot list journal: %v", err)
	}

	pbJournalList := make([]*pb.Journal, 0, len(journalList))
	for _, journal := range journalList {
		pbJournalList = append(pbJournalList, convertJournal(journal))
	}

	return &pb.QueryJournalsResponse{
		Journals: pbJournalList,
	}, nil
}