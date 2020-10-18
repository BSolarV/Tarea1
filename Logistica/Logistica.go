package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
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

//Server (server)
type Server struct {
	registry map[string]*ProtoLogistic.Package

	mutex sync.Mutex

	// Active Queues
	packageCount int

	retailQueue   []*ProtoLogistic.Package
	priorityQueue []*ProtoLogistic.Package
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

//Interacciones con el Cliente

// DeliverPackage hace la acción después del pedido del cliente.
func (s *Server) DeliverPackage(ctx context.Context, clientPackage *ProtoLogistic.Package) (*ProtoLogistic.Package, error) {
	//Se guardan en el registro
	s.mutex.Lock()
	s.packageCount++
	clientPackage.IDPaquete = strconv.Itoa(s.packageCount)
	clientPackage.Seguimiento = strconv.Itoa(s.packageCount)
	clientPackage.Estado = "En bodega"
	if clientPackage.GetTipo() == 1 { // Retail = 1
		clientPackage.Seguimiento = "0"
	}
	s.registry[clientPackage.GetIDPaquete()] = clientPackage

	//Se añaden los objetos a la cola correspondiente
	if clientPackage.GetTipo() == 1 {
		s.retailQueue = append(s.retailQueue, clientPackage)
	} else if clientPackage.GetTipo() == 2 {
		s.priorityQueue = append(s.priorityQueue, clientPackage)
	} else {
		s.normalQueue = append(s.normalQueue, clientPackage)
	}
	s.mutex.Unlock()
	fmt.Print("DeliverPackage")
	fmt.Printf("Id package : %s   Estado:  %s\n", clientPackage.GetIDPaquete(), clientPackage.GetEstado())
	return clientPackage, nil
}

//CheckStatus revisa el estado del paquete. Lo acciona el Cliente
func (s *Server) CheckStatus(ctx context.Context, clientPackage *ProtoLogistic.Package) (*ProtoLogistic.Package, error) {
	//Obtengo codigo de seguimiento
	seguimiento := clientPackage.GetSeguimiento()
	paq := s.registry[seguimiento]

	fmt.Print("CheckStatus")
	fmt.Printf("Id package : %s   Estado:  %s\n", paq.GetIDPaquete(), paq.GetEstado())
	return paq, nil
}

//############################################

//Interacciones con los camiones

//AskPackage es la acción de pedir un paquete. Lo hace el camión
func (s *Server) AskPackage(ctx context.Context, truck *ProtoLogistic.Truck) (*ProtoLogistic.Package, error) {
	tipoCamion := truck.GetType()
	var paq *ProtoLogistic.Package
	if tipoCamion == 1 {
		s.mutex.Lock()
		if s.retailQueue != nil {
			paq = s.retailQueue[0]
			s.retailQueue = s.retailQueue[1:]
		} else if s.priorityQueue != nil {
			paq = s.priorityQueue[0]
			s.priorityQueue = s.priorityQueue[1:]
		}
		s.mutex.Unlock()
	} else {
		s.mutex.Lock()
		if s.priorityQueue != nil {
			paq = s.priorityQueue[0]
			s.priorityQueue = s.priorityQueue[1:]
		} else if s.normalQueue != nil {
			paq = s.normalQueue[0]
			s.normalQueue = s.normalQueue[1:]
		}
		s.mutex.Unlock()
	}
	if paq != nil {
		s.mutex.Lock()
		paq.Estado = "En camino"
		s.mutex.Unlock()
		fmt.Print("CheckStatus")
		fmt.Printf("Id package : %s   Estado:  %s\n", paq.GetIDPaquete(), paq.GetEstado())
	}

	return paq, nil
}

//FinishPackage es la acción que se hace cuando el camión termina la entrega
func (s *Server) FinishPackage(ctx context.Context, truckPackage *ProtoLogistic.Package) (*ProtoLogistic.Empty, error) {
	//Se actualiza el registro
	s.mutex.Lock()
	s.registry[truckPackage.GetIDPaquete()] = truckPackage
	s.mutex.Unlock()

	fmt.Print("FinishStatus")
	fmt.Printf("Id package : %s   Estado:  %s\n", truckPackage.GetIDPaquete(), truckPackage.GetEstado())

	return &ProtoLogistic.Empty{}, nil
}

//############################
