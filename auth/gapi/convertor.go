package gapi

import (
	"remind/auth/pb"
	"remind/auth/token"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertPayload(payload *token.Payload) *pb.Payload {
	return &pb.Payload{
		Id: payload.ID.String(),
		Username: payload.Username,
		IssuedAt: timestamppb.New(payload.IssuedAt),
		ExpiredAt: timestamppb.New(payload.ExpiredAt),
	}
}