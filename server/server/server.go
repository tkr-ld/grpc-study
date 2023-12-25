package main

import (
	grpc_weather "github.com/tkr-ld/grpc-study/server/gen/api/proto"
	"github.com/tkr-ld/grpc-study/server/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listten")
	}

	server := grpc.NewServer()

	grpc_weather.RegisterWeatherServiceServer(server, handler.NewWeatherHandler())

	reflection.Register(server)

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
