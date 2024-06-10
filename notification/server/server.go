package server

import (
	"context"
	"notification/firebase"
	"notification/database"
	"notification/model"

	"notification/pb"
	repository "notification/repository"
     "go.mongodb.org/mongo-driver/bson"

)

type Server struct{
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
    err := database.FCMTokenCollection.FindOne(ctx, bson.M{"user_id": req.GetUserId()}).Decode(&token)
    if err != nil {
        return &pb.MessageResponse{
            Success: false,
            ErrorMessage: err.Error(),
        }, nil
    }


    err = firebase.SendNotification(ctx, token.Token, req.Title, req.Body)
    if err != nil {
        return &pb.MessageResponse{Success: false, ErrorMessage: err.Error()}, nil
    }
    return &pb.MessageResponse{Success: true}, nil
}

