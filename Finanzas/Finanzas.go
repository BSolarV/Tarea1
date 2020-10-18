package main

import (
	"fmt"
	"math"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Finanzas")
	conn, err := amqp.Dial("amqp://winducloveer:secret@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Conected")

}

type Paquete struct {
	idPaquete     string
	descripcion   string
	tipo          string
	intentos      int
	estado        string
	valorOriginal int
	balance       int
}

type Finanzas struct {
	registry []*Paquete
	balance  int
}

func (f *Finanzas) añadirPaquete(pkg *Paquete) {
	pkg.balance = 0
	if pkg.tipo == "Prioritario" {
		pkg.balance += int(math.Round(float64(pkg.valorOriginal) * 0.3))
	}
	if pkg.estado == "Recibido" {
		pkg.balance += pkg.valorOriginal
	} else if pkg.tipo == "Retail" {
		pkg.balance += pkg.valorOriginal
	}
	pkg.balance = pkg.balance - 10*(pkg.intentos-1)

	f.registry = append(f.registry, pkg)
	f.balance += pkg.balance
}
