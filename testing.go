package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/BSolarV/Tarea1/ProtoLogistic"
)

func main() {
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

		printPackage(&packageToAdd)
		result = append(result, packageToAdd)
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
	fmt.Println("Printeado!\n")
}
