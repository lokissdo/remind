package gapi

import (
	"context"
	"remind/auth/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) VerifyToken(ctx context.Context, req *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	payload, err := server.tokenMaker.VerifyToken(req.GetToken())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	rsp := &pb.VerifyTokenResponse{
		Payload: convertPayload(payload),
	}

	return rsp, nil
}