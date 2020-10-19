package main

import (
	"bufio"
	"encoding/csv"
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

	fmt.Print("Máximo tiempo de inactividad (en minutos): ")
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	FinishMargin, err := strconv.Atoi(text)
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
		"WinduCloveerQueue",
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
	finishLine := time.Now().Add(time.Duration(FinishMargin) * time.Minute)

	go func() {
		for {
			if time.Now().After(finishLine) {
				forever <- true
			}
		}
	}()

	var pkg Paquete
	go func() {
		for binaryPkg := range pkgs {
			err := json.Unmarshal(binaryPkg.Body, &pkg)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			finanzas.añadirPaquete(&pkg)
			writeRegistry(&pkg, -999999)
			finishLine = time.Now().Add(time.Duration(FinishMargin) * time.Minute)
		}
	}()

	<-forever
	fmt.Printf("El balance final es: %d", finanzas.balance)
	writeRegistry(&pkg, int(finanzas.balance))
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

// Finanzas : Estructura para mantener registros
type Finanzas struct {
	registry []Paquete
	balance  int
}

func (f *Finanzas) añadirPaquete(pkg *Paquete) {
	pkg.Balance = 0
	if pkg.Tipo == "Prioritario" {
		pkg.Balance += int(math.Round(float64(pkg.ValorOriginal) * 0.3))
	}
	if pkg.Estado == "Recibido" {
		pkg.Balance += pkg.ValorOriginal
	} else if pkg.Tipo == "Retail" {
		pkg.Balance += pkg.ValorOriginal
	}
	pkg.Balance = pkg.Balance - 10*(pkg.Intentos-1)

	f.registry = append(f.registry, *pkg)
	f.balance += pkg.Balance
}

func (f *Finanzas) printRegistry() {
	fmt.Println("-------------------------------------")
	fmt.Printf(" %s | %s | %s | %s | %s | %s | %s |\n",
		leftjust("ID", 3), leftjust("Descripcion", 25), leftjust("Tipo", 10),
		leftjust("Intentos", 8), leftjust("Estado", 15), leftjust("Valor", 6), leftjust("Balance", 8))
	for _, pkg := range f.registry {
		fmt.Printf(" %s | %s | %s | %s | %s | %s | %s |\n",
			leftjust(pkg.IDPaquete, 3), leftjust(pkg.Descripcion, 25), leftjust(pkg.Tipo, 13),
			leftjust(strconv.Itoa(pkg.Intentos), 8), leftjust(pkg.Estado, 15), leftjust(strconv.Itoa(pkg.ValorOriginal), 6),
			leftjust(strconv.Itoa(pkg.Balance), 8))
	}
	fmt.Printf(" %s   %s   %s   %s   %s   %s | %s |\n",
		leftjust(" ", 3), leftjust(" ", 25), leftjust(" ", 10),
		leftjust(" ", 8), leftjust(" ", 15), leftjust("Total:", 6), leftjust(strconv.Itoa(f.balance), 8))
	fmt.Println("-------------------------------------")
}

func leftjust(s string, n int) string {
	if len(s) >= n {
		return s
	}
	return s + strings.Repeat(" ", (n-len(s)))
}

func writeRegistry(pkg *Paquete, balance int) {

	file, err := os.OpenFile("registroFinanzas.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		file, err = os.Create("registroFinanzas.csv")
		if err != nil {
			panic(err)
		}
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	if balance != -999999 {
		err = writer.Write([]string{"total:", strconv.Itoa(balance)})
		if err != nil {
			fmt.Println(err)
		}
	}
	toWrite := []string{
		pkg.IDPaquete,
		pkg.Descripcion,
		pkg.Tipo,
		strconv.Itoa(pkg.Intentos),
		pkg.Estado,
		strconv.Itoa(pkg.ValorOriginal),
		strconv.Itoa(pkg.Balance)}
	fmt.Println("Escribirndo: ", toWrite)
	err = writer.Write(toWrite)
	if err != nil {
		fmt.Println(err)
	}
}
