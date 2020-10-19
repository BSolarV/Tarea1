package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	ProtoLogistic "github.com/BSolarV/Tarea1/ProtoLogistic"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.10.28.63:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect: %s", err)
	}
	defer conn.Close()

	clientService := ProtoLogistic.NewProtoLogisticServiceClient(conn)

	pymesPackages := ParsePymes()
	retailPackages := ParseRetail()

	// For testing
	reader := bufio.NewReader(os.Stdin)
	//Seteando comportamiento del cliente
	fmt.Println("Ingrese nro. tipo cliente:  [Retail : 0 , Pyme : 1]  ")
	tipoCliente, _ := reader.ReadString('\n')
	tipoCliente = strings.Replace(tipoCliente, "\n", "", -1)
	tipoCliente = strings.Replace(tipoCliente, "\r", "", -1)

	fmt.Println("Ingrese el tiempo entre pedidos del cliente: ")
	tiempoEspera, _ := reader.ReadString('\n')
	tiempoEspera = strings.Replace(tiempoEspera, "\n", "", -1)
	tiempoEspera = strings.Replace(tiempoEspera, "\r", "", -1)

	waitTime, _ := strconv.Atoi(tiempoEspera)

	var mutex sync.Mutex

	var SegCodes []string

	auxPaq := retailPackages
	if tipoCliente == "1" {
		auxPaq = pymesPackages
	}

	var wg sync.WaitGroup

	//"0 : Retail ,  1 : Pyme "
	wg.Add(1)
	go func() { //Hago peticiones
		defer wg.Done()

		for len(auxPaq) != 0 {
			mutex.Lock()
			if strings.Compare("0", tipoCliente) == 0 {
				if len(retailPackages) == 0 {
					fmt.Println("We are out of that.")
					continue
				}

				pack := retailPackages[0]
				retailPackages = retailPackages[1:]
				_, err := clientService.DeliverPackage(context.Background(), pack)
				if err != nil {
					panic(err)
				}

			} else if strings.Compare("1", tipoCliente) == 0 {
				if len(pymesPackages) == 0 {
					fmt.Println("We are out of that.")
					continue
				}
				pack := pymesPackages[0]
				pymesPackages = pymesPackages[1:]
				response, err := clientService.DeliverPackage(context.Background(), pack)
				if err != nil {
					panic(err)
				}

				SegCodes = append(SegCodes, response.GetSeguimiento())
				fmt.Printf("El código de seguimiento es: %s \n", response.GetSeguimiento())
			}
			auxPaq = auxPaq[1:]
			mutex.Unlock()
			time.Sleep(time.Duration(waitTime) * time.Second)
		}
	}()
	if tipoCliente == "1" {
		wg.Add(1)
		go func() { //Hago Seguimiento
			defer wg.Done()
			for len(auxPaq) != 0 {
				if len(SegCodes) == 0 {
					continue
				}
				mutex.Lock()
				index := int(rand.Intn(len(SegCodes)))
				packag, err := clientService.CheckStatus(context.Background(), &ProtoLogistic.Package{Seguimiento: SegCodes[index]})
				mutex.Unlock()
				if err != nil {
					panic(err)
				}
				fmt.Println("Consultando Estado")
				fmt.Printf("Id: %s; type: %s; valor: %d; Origen: %s; Destino: %s; \n \t desc: %s \n \t ******* ESTADO : %s ********\n:",
					packag.GetIDPaquete(),
					packag.GetTipo(),
					packag.GetValor(),
					packag.GetOrigen(),
					packag.GetDestino(),
					packag.GetProducto(),
					packag.GetEstado())
				fmt.Println("Printeado!")
				mutex.Lock()
				swap := reflect.Swapper(SegCodes)
				swap(0, index)
				SegCodes = SegCodes[1:]
				mutex.Unlock()
				time.Sleep(time.Duration(waitTime) * time.Second)
			}
		}()
	}
	wg.Wait()

}

// ParsePymes Funcion para parsear CSV Pymes
func ParsePymes() []*ProtoLogistic.Package {

	var result []*ProtoLogistic.Package

	// Open the file
	csvfile, err := os.Open("files/pymes.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			continue
		}
		if record[0] == "id" {
			continue
		}
		var packageType ProtoLogistic.PackageType
		if record[5] == "1" {
			packageType = 2
		} else {
			packageType = 3
		}
		value64, err := strconv.ParseInt(record[2], 10, 32)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		value := int32(value64)
		packageToAdd := ProtoLogistic.Package{
			IDPaquete: record[0],
			Producto:  record[1],
			Tipo:      packageType,
			Valor:     value,
			Origen:    record[3],
			Destino:   record[4]}
		result = append(result, &packageToAdd)
	}
	return result
}

// ParseRetail Funcion para parsear CSV Retail
func ParseRetail() []*ProtoLogistic.Package {

	var result []*ProtoLogistic.Package

	// Open the file
	csvfile, err := os.Open("files/retail.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			continue
		}
		if record[0] == "id" {
			continue
		}
		var packageType ProtoLogistic.PackageType
		packageType = 1
		value64, err := strconv.ParseInt(record[2], 10, 32)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		value := int32(value64)
		packageToAdd := ProtoLogistic.Package{
			IDPaquete: record[0],
			Producto:  record[1],
			Tipo:      packageType,
			Valor:     value,
			Origen:    record[3],
			Destino:   record[4]}
		result = append(result, &packageToAdd)
	}
	return result
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
