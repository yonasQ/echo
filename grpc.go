package main

import (
	"context"
	"log"
	"net"

	grpc_rest_echo "echo/github.com/yonasQ/grpc-rest-echo"

	"google.golang.org/grpc"
)

type server struct {
	grpc_rest_echo.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, in *grpc_rest_echo.EchoRequest) (*grpc_rest_echo.EchoResponse, error) {
	return &grpc_rest_echo.EchoResponse{Message: in.Message}, nil
}

func startGRPCServer(lis net.Listener) {
	s := grpc.NewServer()
	grpc_rest_echo.RegisterEchoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
