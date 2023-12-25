package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

)
func main() {
	lis, err := net.Listen("tcp","50051")
	if err != nil {
		log.Fatalf("failed to listten")
	}

	server := grpc
}