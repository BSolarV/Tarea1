package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/streadway/amqp"

	ProtoLogistic "github.com/BSolarV/Tarea1/ProtoLogistic"
	"google.golang.org/grpc"
)

var finishLine time.Time
var FinishMargin int

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Máximo tiempo de inactividad (en minutos): ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	_FinishMargin, err := strconv.Atoi(text)
	FinishMargin = _FinishMargin
	if err != nil {
		panic(err)
	}

	//Conexión con grpc
	lis, err := net.Listen("tcp", "10.10.28.63:9000")

	if err != nil {
		log.Fatalf("Fail listening on port 9000: %v", err)
	}

	//Conexión al servidor de rabbitmq
	con, er := amqp.Dial("amqp://winducloveer:secret@10.10.28.66:5672/")
	if er != nil {
		fmt.Println(er)
		panic(er)
	}
	defer con.Close()

	ch, er := con.Channel()
	if er != nil {
		fmt.Println(er)
		panic(er)
	}
	defer ch.Close()

	q, er := ch.QueueDeclare(
		"WinduCloveerQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if er != nil {
		fmt.Println(er)
		panic(er)
	}

	srv := NewServer(con, ch, &q)

	grpcServer := grpc.NewServer()

	ProtoLogistic.RegisterProtoLogisticServiceServer(grpcServer, srv)

	finishLine = time.Now().Add(time.Duration(FinishMargin) * time.Minute)

	go func() {
		for {
			if time.Now().After(finishLine) {
				fmt.Printf("cerrando porque %s after %s", time.Now().Format("2006.01.02 15:04:05"), finishLine.Format("2006.01.02 15:04:05"))
				ch.Close()
				con.Close()
				grpcServer.Stop()
				lis.Close()
				break
			}
		}
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to mount GRPC server on port 9000: %v", err)
	}

}

//Server (server)
type Server struct {
	//Registro de Odín
	registry       map[string]*ProtoLogistic.Package
	registryFile   *os.File
	registryWriter *csv.Writer

	//Para locks y unlocks
	mutex sync.Mutex

	//Atributos conexión rabbitmq
	con *amqp.Connection
	ch  *amqp.Channel
	q   *amqp.Queue

	// Active Queues
	packageCount int

	retailQueue   []*ProtoLogistic.Package
	priorityQueue []*ProtoLogistic.Package
	normalQueue   []*ProtoLogistic.Package
	// Punteros a los paquetes del registro, para facilitar la modificacion de registro

}

//NewServer es el constructor del Server
func NewServer(con *amqp.Connection, ch *amqp.Channel, q *amqp.Queue) *Server {
	var srv Server
	srv.registry = make(map[string]*ProtoLogistic.Package)

	srv.con = con
	srv.ch = ch
	srv.q = q

	return &srv
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
	fmt.Println(time.Now().Format("2006.01.02 15:04:05"), finishLine.Format("2006.01.02 15:04:05"))
	finishLine = time.Now().Add(time.Duration(FinishMargin) * time.Minute)
	fmt.Println(time.Now().Format("2006.01.02 15:04:05"), finishLine.Format("2006.01.02 15:04:05"))
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

	writeRegistry(clientPackage)

	//Se añaden los objetos a la cola correspondiente
	if clientPackage.GetTipo() == 1 {
		s.retailQueue = append(s.retailQueue, clientPackage)
	} else if clientPackage.GetTipo() == 2 {
		s.priorityQueue = append(s.priorityQueue, clientPackage)
	} else {
		s.normalQueue = append(s.normalQueue, clientPackage)
	}
	s.mutex.Unlock()
	fmt.Print("DeliverPackage: ")
	fmt.Printf("Id package : %s   Estado:  %s\n", clientPackage.GetIDPaquete(), clientPackage.GetEstado())
	return clientPackage, nil
}

//CheckStatus revisa el estado del paquete. Lo acciona el Cliente
func (s *Server) CheckStatus(ctx context.Context, clientPackage *ProtoLogistic.Package) (*ProtoLogistic.Package, error) {
	//Obtengo codigo de seguimiento
	seguimiento := clientPackage.GetSeguimiento()
	paq := s.registry[seguimiento]

	fmt.Print("CheckStatus: ")
	fmt.Printf("Id package : %s   Estado:  %s\n", paq.GetIDPaquete(), paq.GetEstado())
	return paq, nil
}

//############################################

//Interacciones con los camiones

//AskPackage es la acción de pedir un paquete. Lo hace el camión
func (s *Server) AskPackage(ctx context.Context, truck *ProtoLogistic.Truck) (*ProtoLogistic.Package, error) {
	tipoCamion := truck.GetType()
	var paq *ProtoLogistic.Package
	paq = &ProtoLogistic.Package{IDPaquete: "-1"}
	if tipoCamion == 1 {
		s.mutex.Lock()
		if len(s.retailQueue) != 0 {
			paq = s.retailQueue[0]
			s.retailQueue = s.retailQueue[1:]
		} else if len(s.priorityQueue) != 0 {
			paq = s.priorityQueue[0]
			s.priorityQueue = s.priorityQueue[1:]
		}
		s.mutex.Unlock()
	} else {
		s.mutex.Lock()
		if len(s.priorityQueue) != 0 {
			paq = s.priorityQueue[0]
			s.priorityQueue = s.priorityQueue[1:]
		} else if len(s.normalQueue) != 0 {
			paq = s.normalQueue[0]
			s.normalQueue = s.normalQueue[1:]
		}
		s.mutex.Unlock()
	}
	if paq.GetIDPaquete() != "-1" {
		s.mutex.Lock()
		paq.Estado = "En camino"
		s.mutex.Unlock()
		fmt.Print("AskPackage: ")
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

	fmt.Print("FinishPackage: ")
	fmt.Printf("Id package : %s   Estado:  %s\n", truckPackage.GetIDPaquete(), truckPackage.GetEstado())

	s.mutex.Lock()
	auxPaq := paqueteFinanza(truckPackage)

	s.SendToFinanzas(auxPaq)
	s.mutex.Unlock()

	return &ProtoLogistic.Empty{}, nil
}

//############################

//Interacción con Finanzas

//SendToFinanzas es la función que envia los paquetes a la cola de finanzas
func (s *Server) SendToFinanzas(pkg Paquete) {
	body, er := json.Marshal(pkg)
	if er != nil {
		fmt.Println(er)
		panic(er)
	}
	er = s.ch.Publish(
		"",
		"WinduCloveerQueue",
		false,
		false,
		amqp.Publishing{ContentType: "application/json",
			Body: body,
		},
	)
	if er != nil {
		fmt.Println(er)
		panic(er)
	}

}

func printPackage(packag *ProtoLogistic.Package) {
	fmt.Println("Printeando Paquete")
	fmt.Printf("Id: %s; type: %s; valor: %d; Origen: %s; Destino: %s; \n desc: %s \n",
		packag.GetIDPaquete(),
		packag.GetTipo(),
		packag.GetValor(),
		packag.GetOrigen(),
		packag.GetDestino(),
		packag.GetProducto())
	fmt.Println("Printeado!")
}

// Paquete : Estructura para facilitar marshaling en Json
type Paquete struct {
	IDPaquete     string `json:"idPaquete"`
	Descripcion   string `json:"descripcion"`
	Tipo          string `json:"tipo"`
	Estado        string `json:"estado"`
	Intentos      int    `json:"intentos"`
	ValorOriginal int    `json:"valorOriginal"`
	Balance       int    `json:"balance"`
}

func paqueteFinanza(pkg *ProtoLogistic.Package) Paquete {
	var paq Paquete
	paq.IDPaquete = pkg.GetIDPaquete()
	paq.Descripcion = pkg.GetProducto()
	tipo := pkg.GetTipo()
	if tipo == 1 {
		paq.Tipo = "Retail"
	} else if tipo == 2 {
		paq.Tipo = "Prioritario"
	} else {
		paq.Tipo = "Normal"
	}
	paq.Estado = pkg.GetEstado()
	paq.Intentos = int(pkg.GetIntentos())
	paq.ValorOriginal = int(pkg.GetValor())
	return paq
}

func writeRegistry(clientPackage *ProtoLogistic.Package) {
	file, err := os.OpenFile("registroLogistica.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		file, err = os.Create("registroLogistica.csv")
		if err != nil {
			panic(err)
		}
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escritura en Registro
	var tipo string
	if clientPackage.GetTipo() == 1 {
		tipo = "retail"
	} else if clientPackage.GetTipo() == 2 {
		tipo = "prioritario"
	} else {
		tipo = "normal"
	}
	toWrite := []string{
		time.Now().Format("2006.01.02 15:04:05"),
		clientPackage.GetIDPaquete(),
		tipo,
		clientPackage.GetProducto(),
		strconv.Itoa(int(clientPackage.GetValor())),
		clientPackage.GetOrigen(),
		clientPackage.GetDestino(),
		clientPackage.GetIDPaquete()}
	fmt.Println("Escribirndo: ", toWrite)
	err = writer.Write(toWrite)
	if err != nil {
		fmt.Println(err)
	}
}
