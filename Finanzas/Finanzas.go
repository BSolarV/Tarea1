package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Cada cuanto mostrar Registro (en segundos): ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	WaitRegistro, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}

	fmt.Println("Finanzas")
	conn, err := amqp.Dial("amqp://winducloveer:secret@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Conected")

	channel, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer channel.Close()

	fmt.Println("Iniciando Sistema Financiero")

	finanzas := Finanzas{}

	pkgs, err := channel.Consume(
		"WinducloveerQueue",
		"",
		true,
		false,
		false,
		false,
		nil)

	go func() {
		for {
			time.Sleep(time.Duration(WaitRegistro) * time.Second)
			finanzas.printRegistry()
		}
	}()

	forever := make(chan bool)
	var pkg *Paquete
	go func() {
		for binaryPkg := range pkgs {
			fmt.Println("Paquete recibido")
			err := json.Unmarshal(binaryPkg.Body, pkg)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			finanzas.añadirPaquete(pkg)
		}
	}()
	<-forever
}

// Paquete : Estructura para facilitar marshaling en Json
type Paquete struct {
	idPaquete     string
	descripcion   string
	tipo          string
	intentos      int
	estado        string
	valorOriginal int
	balance       int
}

// Finanzas : Estructura para mantener registros
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

func (f *Finanzas) printRegistry() {
	fmt.Printf(" %s | %s | %s | %s | %s | %s | %s |\n",
		leftjust("ID", 5), leftjust("Descripcion", 15), leftjust("Tipo", 10),
		leftjust("Intentos", 8), leftjust("Estado", 15), leftjust("Valor", 6), leftjust("Balance", 8))
	for _, pkg := range f.registry {
		fmt.Printf(" %s | %s | %s | %s | %s | %s | %s |\n",
			leftjust(pkg.idPaquete, 5), leftjust(pkg.descripcion, 15), leftjust(pkg.tipo, 10),
			leftjust(strconv.Itoa(pkg.intentos), 8), leftjust(pkg.estado, 15), leftjust(strconv.Itoa(pkg.valorOriginal), 6),
			leftjust(strconv.Itoa(pkg.balance), 8))
	}
	fmt.Printf(" %s   %s   %s   %s   %s   %s | %s |\n",
		leftjust(" ", 5), leftjust(" ", 15), leftjust(" ", 10),
		leftjust(" ", 8), leftjust(" ", 15), leftjust(" ", 6), leftjust(strconv.Itoa(f.balance), 8))
}

func leftjust(s string, n int) string {
	if len(s) >= n {
		return s
	}
	return s + strings.Repeat(" ", (n-len(s)))
}
