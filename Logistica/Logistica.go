package main

import (
	"context"
	"log"
	"net"

	ProtoLogistic "github.com/BSolarV/Tarea1/ProtoLogistic"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Fail listening on port 9000: %v", err)
	}

	srv := Server{}

	grpcServer := grpc.NewServer()

	ProtoLogistic.RegisterProtoLogisticServiceServer(grpcServer, &srv)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to mount GRPC server on port 9000: %v", err)
	}

}

type Server struct {
	registry []ProtoLogistic.Package

	// Active Queues
	packageCount  int
	retailQueue   []*ProtoLogistic.Package
	priotiryQueue []*ProtoLogistic.Package
	normalQueue   []*ProtoLogistic.Package
	// Punteros a los paquetes del registro, para facilitar la modificacion de registro

}

/*
    rpc DeliverPackage(Package) returns (Empty) {}
    rpc CheckStatus(Package) returns (Package) {}

    rpc AskPackage(Truck) returns (Package) {}
	rpc FinishPackage(Package) returns (Empty) {}
*/

func (s *Server) DeliverPackage(ctx context.Context, clientPackage *ProtoLogistic.Package) (*ProtoLogistic.Empty, error) {
	return nil, nil
}

func (s *Server) CheckStatus(ctx context.Context, clientPackage *ProtoLogistic.Package) (*ProtoLogistic.Package, error) {
	return nil, nil
}

func (s *Server) AskPackage(ctx context.Context, truck *ProtoLogistic.Truck) (*ProtoLogistic.Package, error) {
	return nil, nil
}

func (s *Server) FinishPackage(ctx context.Context, truckPackage *ProtoLogistic.Package) (*ProtoLogistic.Empty, error) {
	return nil, nil
}
