package main

import (
	"context"
	grpc_rest_echo "echo/github.com/yonasQ/grpc-rest-echo"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func startHTTPServer(lis net.Listener) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	ctx := context.Background()

	if err := grpc_rest_echo.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, lis.Addr().String(), opts); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}

	if err := http.Serve(lis, mux); err != nil {
		log.Fatalf("Failed to serve HTTP server: %v", err)
	}
}
