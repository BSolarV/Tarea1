package main

import (
	"context"
	"log"
	"net"
	"sync"

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
	registry map[ string ]ProtoLogistic.Package

	mutex sync.Mutex

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

/*
Codigo de Getters
	idPaquete := clientPackage.GetIDPaquete()
	tipo := clientPackage.GetTipo()
	valor := clientPackage.GetValor()
	origen := clientPackage.GetOrigen()
	destino := clientPackage.GetDestino()
	intentos := clientPackage.GetIntentos()
	estado := clientPackage.GetEstado()
	seguimiento := clientPackage.GetSeguimiento()
*/

// Interacciones con el Cliente
func (s *Server) DeliverPackage(ctx context.Context, clientPackage *ProtoLogistic.Package) (*ProtoLogistic.Empty, error) {
	//Se guardan en el registro
	s.mutex.Lock()
	s.packageCount += 1
	clientPackage.IDPaquete = s.packageCount
	clientPackage.Seguimiento = s.packageCount
	if clientPackage.GetTipo == 1{ // Retail = 1
		clientPackage.Seguimiento = 0
	}
	s.registry[clientPackage.GetIDPaquete()] = *clientPackage
	s.mutex.Unlock()

	return ProtoLogistic.Empty{}, nil
}

func (s *Server) CheckStatus(ctx context.Context, clientPackage *ProtoLogistic.Package) (*ProtoLogistic.Package, error) {
	//Obtengo codigo de seguimiento
	seguimiento := clientPackage.GetSeguimiento()
	paq := s.registry[seguimiento]

	return &paq, nil
}

//############################################

//Interacciones con los camiones

func (s *Server) AskPackage(ctx context.Context, truck *ProtoLogistic.Truck) (*ProtoLogistic.Package, error) {
	tipoCamion := truck.GetType()

	if tipoCamion == 1 {
		for _,id := range s.registry{
			if s.registry[id].Estado == "En bodega" && s.registry[id].Tipo == 1 {
				s.mutex.Lock()
				s.registry[id].Estado = "En camino"
				s.mutex.Unlock()
				return nil, nil
			} 
		}
		for _,id := range s.registry{
			if s.registry[id].Estado == "En bodega" && s.registry[id].Tipo == 2 {
				s.mutex.Lock()
				s.registry[id].Estado = "En camino"
				s.mutex.Unlock()
				return nil, nil
			} 
		}
		return nil,nil
	}
	else{
		for _,id := range s.registry{
			if s.registry[id].Estado == "En bodega" && s.registry[id].Tipo == 1 {
				s.mutex.Lock()
				s.registry[id].Estado = "En camino"
				s.mutex.Unlock()
				return nil, nil
			} 
		}
		for _,id := range s.registry{
			if s.registry[id].Estado == "En bodega" && s.registry[id].Tipo == 2 {
				s.mutex.Lock()
				s.registry[id].Estado = "En camino"
				s.mutex.Unlock()
				return nil, nil
			} 
		}
		return nil,nil
	}
	return nil, nil
}

func (s *Server) FinishPackage(ctx context.Context, truckPackage *ProtoLogistic.Package) (*ProtoLogistic.Empty, error) {
	//Se actualiza el registro
	s.mutex.Lock()
	s.registry[truckPackage.GetIDPaquete].Estado = *truckPackage
	s.mutex.Unlock()

	return ProtoLogistic.Empty{}, nil
}

//############################
