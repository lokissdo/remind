package main

import (
	"context"
	"errors"
	"os"

	"fmt"
	"net/http"

	"remind/gateway/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"remind/gateway/middleware"

	"remind/pkg/logger"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authgen "remind/auth/pb"
	todogen "remind/todo/pb"

	usergen "remind/user/pb"
	journalgen "remind/journal/pb"
	aigen "remind/ai/pb"

)

// gRPC Gateway
func newGateway(
	ctx context.Context,
	cfg *config.Config,
	opts []gwruntime.ServeMuxOption,
) (http.Handler, error) {
	mux := gwruntime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024 * 1024 * 1024)),
	}

	userEndpoint := fmt.Sprintf("%s:%d", cfg.UserHost, cfg.UserPort)
	err := usergen.RegisterUserRemindHandlerFromEndpoint(ctx, mux, userEndpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	authEndpoint := fmt.Sprintf("%s:%d", cfg.AuthHost, cfg.AuthPort)
	err = authgen.RegisterAuthRemindHandlerFromEndpoint(ctx, mux, authEndpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	journalEndpoint := fmt.Sprintf("%s:%d", cfg.JournalHost, cfg.JournalPort)
	err = journalgen.RegisterJournalRemindHandlerFromEndpoint(ctx, mux, journalEndpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	todoEndpoint := fmt.Sprintf("%s:%d", cfg.TodoHost, cfg.TodoPort)
	err = todogen.RegisterTodoServiceHandlerFromEndpoint(ctx, mux, todoEndpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	aiEndpoint := fmt.Sprintf("%s:%d", cfg.AIHost, cfg.AIPort)
	err = aigen.RegisterAIServiceHandlerFromEndpoint(ctx, mux, aiEndpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	if cfg.Environment.Name == "dev" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	mux := http.NewServeMux()

	gw, err := newGateway(ctx, cfg, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create gateway")
	}

	mux.Handle("/", gw)

	securityConfigs, err := middleware.LoadAPISecurityConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load security configs")
	}

	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler: logger.HttpLogger(middleware.AllowCORS(middleware.AuthMiddleWare(mux, securityConfigs))),
	}

	go func() {
		<-ctx.Done()
		log.Info().Msg("shutting down the http server")

		if err := s.Shutdown(context.Background()); err != nil {
			log.Fatal().Err(err).Msg("failed to shutdown http server")
		}
	}()
	
	log.Info().Msgf("start listening... address: %s:%d", cfg.Host, cfg.Port)

	err = s.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal().Err(err).Msg("failed to listen and serve")
	} else if err != nil {
		log.Fatal().Err(err).Msg("Unexpected error")
	}
}
