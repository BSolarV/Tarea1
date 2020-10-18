package main

import (
	"fmt"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"sync"

	ProtoLogistic "github.com/BSolarV/Tarea1/ProtoLogistic"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect: %s", err)
	}
	defer conn.Close()

	clientService := ProtoLogistic.NewProtoLogisticServiceClient(conn)

	pymesPackages := ParsePymes()
	retailPackages := ParseRetail()

	var wg sync.WaitGroup

	// For testing
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("0 : Retail - 1 : Pyme - 3 : codigo -> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("1", text) == 0 {
			if len(pymesPackages) == 0:
				fmt.Println("We are out of that.")
				continue
			package := pymesPackages[0]
			pymesPackages = pymesPackages[1:]
			response , err := clientService.DeliverPackage(context.Background(), &package)
			if err != nil{
				panic(err)
			}
			
		}
	}
	
}

func ParsePymes() []ProtoLogistic.Package {

	var result []ProtoLogistic.Package

	// Open the file
	csvfile, err := os.Open("pymes.csv")
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
		result = append(result, packageToAdd)
	}
	return result
}

func ParseRetail() []ProtoLogistic.Package {

	var result []ProtoLogistic.Package

	// Open the file
	csvfile, err := os.Open("retail.csv")
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
		result = append(result, packageToAdd)
	}
	return result
}
