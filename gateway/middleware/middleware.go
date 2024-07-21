package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "remind/auth/pb"
)

func AuthMiddleWare(next http.Handler, protectedPaths []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg(r.URL.Path)
		for _, path := range protectedPaths {
			if strings.HasPrefix(r.URL.Path, path) {
				token, err := extractToken(r.Header.Get("Authorization"))
				if err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}
				// Validate the token
				username, err := verifyToken(token)
				if err != nil {
					// If the token is invalid, return an unauthorized status
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				// Set the user ID in the request context
				ctx := context.WithValue(r.Context(), "userID", username)
				r = r.WithContext(ctx)
				break
			}
		}
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func verifyToken(token string) (username string, err error) {
	log.Info().Msg("verifying token")
	conn, err := grpc.NewClient("localhost:5003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Err(err).Msg("failed to dial server")
		return
	}
	defer conn.Close()

	c := pb.NewAuthRemindClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rsp, err := c.VerifyToken(ctx, &pb.VerifyTokenRequest{Token: token})
	if err != nil {
		log.Error().Err(err).Msg("failed to verify token")
		return "", err
	}

	log.Info().Msg(rsp.Payload.Username)

	return rsp.GetPayload().GetUsername(), nil
}

func extractToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is not provided")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization header format")
	}
	
	return parts[1], nil
}