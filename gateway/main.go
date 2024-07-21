package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"remind/gateway/config"
	"remind/pkg/logger"

	"log/slog"

	"remind/gateway/handlers"
	"remind/gateway/middleware"

	"github.com/golang/glog"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authgen "remind/auth/pb"
	usergen "remind/user/pb"
)

func newGateway(
	ctx context.Context,
	cfg *config.Config,
	opts []gwruntime.ServeMuxOption,
) (http.Handler, error) {
	mux := gwruntime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

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

	return mux, nil
}

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		glog.Fatalf("Config error: %s", err)
	}

	// set up logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logger.ConvertLogLevel(cfg.Log.Level))

	// integrate Logrus with the slog logger
	slog.New(logger.NewLogrusHandler(logrus.StandardLogger()))

	mux := http.NewServeMux()

	gw, err := newGateway(ctx, cfg, nil)
	if err != nil {
		slog.Error("failed to create a new gateway", err)
	}

	mux.Handle("/", gw)

	protectedPaths := []string{"/v1/update_user"}

	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler: handlers.LogRequestBody(handlers.AllowCORS(middleware.AuthMiddleWare(mux, protectedPaths))),
	}

	go func() {
		<-ctx.Done()
		slog.Info("shutting down the http server")

		if err := s.Shutdown(context.Background()); err != nil {
			slog.Error("failed to shutdown http server", err)
		}
	}()

	slog.Info("start listening...", "address", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))

	err = s.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to listen and serve", err)
	} else if err != nil {
		slog.Error("unexpected error from ListenAndServe", err)
	}
}
