package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	// "google.golang.org/grpc/metadata"

	pb "remind/auth/pb"
)

// AllowCORS allows Cross Origin Resource Sharing from any origin.
// Don't do this without consideration in production systems.
func AllowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}


// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	log.Info().Msgf("Preflight request for %s", r.URL.Path)
}


// healthzServer returns a simple health handler which returns ok.
func HealthzServer(conn *grpc.ClientConn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		if s := conn.GetState(); s != connectivity.Ready {
			http.Error(w, fmt.Sprintf("grpc server is %s", s), http.StatusBadGateway)
			return
		}
		fmt.Fprintln(w, "ok")
	}
}

func AuthMiddleWare(next http.Handler, securityConfigs map[string]SecurityConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config, exists := securityConfigs[r.URL.Path]
		if exists && config.AuthRequired {
			var username string
			for _, method := range config.AuthMethods {
				switch method {
					case "bearer":
						token, err := extractToken(r.Header.Get("Authorization"))
						if err != nil {
							log.Error().Err(err).Msg("failed to extract token")
							http.Error(w, "failed to extract token", http.StatusUnauthorized)
							return
						}
						
						username, err = verifyToken(token)
						if err != nil {
							log.Error().Err(err).Msg("failed to verify token")
							http.Error(w, "failed to verify token", http.StatusUnauthorized)
							return
						}
				}
				if username != "" {
					break
				}
			}
			if username == "" {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			
			r.Header.Set("Grpc-Metadata-username", username)
		}
		next.ServeHTTP(w, r)
	})
}


func verifyToken(token string) (username string, err error) {
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