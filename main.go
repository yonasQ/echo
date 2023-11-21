package main

import (
	"log"
	"net"

	"github.com/soheilhy/cmux"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	m := cmux.New(lis)

	grpcL := m.Match(cmux.HTTP2())
	httpL := m.Match(cmux.HTTP1())

	go startGRPCServer(grpcL)
	go startHTTPServer(httpL)

	log.Println("server is running")

	if err := m.Serve(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
