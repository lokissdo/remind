package main

import (
	"context"
	"fmt"
	"net/http"
	"remind/cmd/proxy/config"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gen "remind/auth/pb"
)

func newGateway(
	ctx context.Context,
	cfg *config.Config,
	opts []gwruntime.ServeMuxOption,
) (http.Handler, error) {
	// productEndpoint := fmt.Sprintf("%s:%d", cfg.ProductHost, cfg.ProductPort)
	// counterEndpoint := fmt.Sprintf("%s:%d", cfg.CounterHost, cfg.CounterPort)
	authEndpoint := fmt.Sprintf("%s:%d", cfg.AuthHost, cfg.AuthPort)

	mux := gwruntime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := gen.RegisterAuthRemindHandlerFromEndpoint(mux, dialOpts...) //

	// err := gen.RegisterProductServiceHandlerFromEndpoint(ctx, mux, productEndpoint, dialOpts)
	// if err != nil {
	// 	return nil, err
	// }

	// err = gen.RegisterCounterServiceHandlerFromEndpoint(ctx, mux, counterEndpoint, dialOpts)
	// if err != nil {
	// 	return nil, err
	// }

	return mux, nil
}
