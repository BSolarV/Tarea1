package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/BSolarV/Tarea1/chat"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Fail listening on port 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to mount GRPC server on port 9000: %v", err)
	}

}
