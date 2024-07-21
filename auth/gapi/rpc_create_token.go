package gapi

import (
	"context"
	"remind/auth/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateToken(ctx context.Context, req *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
	token, payload, err := server.tokenMaker.CreateToken(
		req.GetUsername(), 
		server.config.AccessTokenDuration,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create token: %v", err)
	}

	rsp := &pb.CreateTokenResponse{
		Token: token,
		Payload: convertPayload(payload),
	}

	return rsp, nil
}