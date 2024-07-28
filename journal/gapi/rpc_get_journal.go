package gapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	db "remind/journal/db/sqlc"
	pb "remind/journal/pb"

	"github.com/rs/zerolog/log"
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
		Images:  pbImageList,
		Audios:  pbAudioList,
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

	var journalList []db.Journal
	var err error

	if req.GetSearch() == "" {
		arg := db.GetJournalFromUserParams{
			Username: req.GetUsername(),
			Limit:    limit,
			Offset:   req.GetOffset(),
		}
		journalList, err = server.store.GetJournalFromUser(ctx, arg)
		if err != nil {
			return nil, fmt.Errorf("cannot list journal: %v", err)
		}
	} else {
		arg := db.QueryJournalParams{
			Username:  req.GetUsername(),
			ToTsquery: req.GetSearch(),
			Limit:     limit,
			Offset:    req.GetOffset(),
		}
		journalList, err = server.store.QueryJournal(ctx, arg)
		if err != nil {
			return nil, fmt.Errorf("cannot list journal: %v", err)
		}

	}

	pbJournalList := make([]*pb.Journal, 0, len(journalList))
	for _, journal := range journalList {
		pbJournalList = append(pbJournalList, convertJournal(journal))
	}

	pbImageList := make([]*pb.Image, 0)
	for _, journal := range journalList {
		iamgeList, err := server.store.GetImageOfJournal(ctx, journal.ID)
		if err != nil {
			return nil, fmt.Errorf("cannot get image list: %v", err)
		}
		for _, image := range iamgeList {
			pbImageList = append(pbImageList, convertImage(image))
		}
	}

	return &pb.QueryJournalsResponse{
		Journals: pbJournalList,
		Images:   pbImageList,
	}, nil
}

type Response struct {
	Images   []int64 `json:"images"`
	Journals []int64 `json:"journals"`
}

func (server *Server) AdvancedQueryJournals(ctx context.Context, req *pb.QueryJournalsRequest) (*pb.AdvancedQueryJournalsResponse, error) {
	// if err := server.authorization(ctx, req.GetUsername()); err != nil {
	// 	return nil, fmt.Errorf("authorization failed: %v", err)
	// }

	limit := req.GetLimit()
	if limit == 0 {
		limit = 10
	}

	params := url.Values{}
	params.Add("username", req.GetUsername())
	params.Add("text", req.GetSearch())
	params.Add("limit", fmt.Sprintf("%d", limit))

	baseURL := "http://localhost:7777/api/v1/query"
	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("cannot get response: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cannot get response: %s", body)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal response: %v", err)
	}

	log.Info().Msgf("response: %v", body)

	return &pb.AdvancedQueryJournalsResponse{
		ImageId:  response.Images,
		JournalId: response.Journals,
	}, nil
}
