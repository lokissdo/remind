package server

import (
	"context"
	"remind/notification/database"
	"remind/notification/firebase"
	"remind/notification/model"
	"strings"

	"remind/notification/pb"
	repository "remind/notification/repository"

	"go.mongodb.org/mongo-driver/bson"
)

type Server struct {
	pb.UnimplementedNotificationServiceServer
}

func (s *Server) AddOrUpdateToken(ctx context.Context, req *pb.TokenRequest) (*pb.TokenResponse, error) {
	err := repository.AddOrUpdateFCMToken(ctx, req.UserId, req.Token)
	if err != nil {
		return &pb.TokenResponse{Success: false, ErrorMessage: err.Error()}, nil
	}
	return &pb.TokenResponse{Success: true}, nil
}

func (s *Server) SendMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {

	var token model.FCMToken
	userID := req.GetUserId()
	// remove start and end spaces in userID
	userID = strings.TrimSpace(userID)

	err := database.FCMTokenCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&token)
	if err != nil {
		println("Error: ", err.Error())
		return &pb.MessageResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	err = firebase.SendNotification(ctx, token.Token, req.Title, req.Body)
	if err != nil {
		return &pb.MessageResponse{Success: false, ErrorMessage: err.Error()}, nil
	}
	return &pb.MessageResponse{Success: true}, nil
}
