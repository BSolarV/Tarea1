package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	ProtoLogistic "github.com/BSolarV/Tarea1/ProtoLogistic"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.10.28.63:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect: %s", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Tiempo espera segundo paquete (en segundos): ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	MaxWait, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}

	fmt.Print("Tiempo de viaje (en segundos): ")
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	TravelTime, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)

	truckService := ProtoLogistic.NewProtoLogisticServiceClient(conn)

	for i := 0; i < 3; i++ {
		go func(i int) {

			var truckType ProtoLogistic.TruckType
			if i%2 == 0 {
				truckType = 1
			} else {
				truckType = 2
			}
			truck := Truck{Type: truckType}

			// Debuging
			log.Printf("Camión %d : Iniciando camion del tipo %s\n", i, truckType)

			for {

				// Debuging
				log.Printf("Camión %d : Esperando Primer Paquete ...\n", i)

				for {
					pkg, err := truckService.AskPackage(context.Background(), &ProtoLogistic.Truck{Type: truck.Type})
					if err != nil {
						forever <- true
					}
					if pkg.GetIDPaquete() != "-1" {
						truck.pkgs = append(truck.pkgs, pkg)
						break
					}
				}

				actualTime := time.Now()
				finishTime := actualTime.Add(time.Duration(MaxWait) * time.Second)

				// Debuging
				log.Printf("Camión %d : Esperando Segundo Paquete...\n", i)

				for time.Now().Before(finishTime) {
					pkg, err := truckService.AskPackage(context.Background(), &ProtoLogistic.Truck{Type: truck.Type})
					if err != nil {
						forever <- true
					}
					if pkg.GetIDPaquete() != "-1" {
						truck.pkgs = append(truck.pkgs, pkg)
						break
					}
				}

				if len(truck.pkgs) == 2 {
					if truck.pkgs[0].GetValor() < truck.pkgs[1].GetValor() {
						truck.pkgs[0], truck.pkgs[1] = truck.pkgs[1], truck.pkgs[0]
					}
				}

				for _, pkg := range truck.pkgs {
					var pkgMaxTries int
					if pkg.GetTipo() == 1 {
						pkgMaxTries = 3
					} else if pkg.GetTipo() == 2 {
						estimatedValue := 0.8*float64(pkg.GetValor()) + 0.2*0.3*float64(pkg.GetValor())
						if estimatedValue > 20 {
							pkgMaxTries = 3
						} else if estimatedValue > 10 {
							pkgMaxTries = 2
						} else {
							pkgMaxTries = 1
						}
					} else {
						estimatedValue := 0.8 * float64(pkg.GetValor())
						if estimatedValue > 20 {
							pkgMaxTries = 3
						} else if estimatedValue > 10 {
							pkgMaxTries = 2
						} else {
							pkgMaxTries = 1
						}
					}
					truck.pkgsToDeliver = append(truck.pkgsToDeliver, &pkgOnDeliver{
						pkg:      pkg,
						maxTries: pkgMaxTries})
				}

				// Debuging
				log.Printf("Camión %d : Iniciando viajes\n", i)

				for len(truck.pkgsToDeliver) != 0 {

					truck.pkgsToDeliver[0].addATry()

					time.Sleep(time.Duration(TravelTime) * time.Second)

					if rand.Intn(100) < 80 {

						// Debuging
						log.Printf("Camión %d : Entrega exitosa de %s\n", i, truck.pkgsToDeliver[0].pkg.GetIDPaquete())

						truck.pkgsToDeliver[0].setStatus("Recibido")
						truck.pkgsToDeliver[0].setDeliveredDate(true)
						truck.pkgsDone = append(truck.pkgsDone, truck.pkgsToDeliver[0])
						truck.pkgsToDeliver = truck.pkgsToDeliver[1:]

					} else if truck.pkgsToDeliver[0].checkMaxTries() {

						// Debuging
						log.Printf("Camión %d : Quiting de %s\n", i, truck.pkgsToDeliver[0].pkg.GetIDPaquete())

						truck.pkgsToDeliver[0].setStatus("No Recibido")
						truck.pkgsToDeliver[0].setDeliveredDate(false)
						truck.pkgsDone = append(truck.pkgsDone, truck.pkgsToDeliver[0])
						truck.pkgsToDeliver = truck.pkgsToDeliver[1:]

					} else if len(truck.pkgsToDeliver) == 2 {
						truck.pkgsToDeliver[0], truck.pkgsToDeliver[1] = truck.pkgsToDeliver[1], truck.pkgsToDeliver[0]
					}
				}

				time.Sleep(time.Duration(TravelTime) * time.Second)

				// Debuging
				log.Printf("Camión %d : De vuelta en central.\n", i)

				for _, pkg := range truck.pkgsDone {
					truck.registry = append(truck.registry, pkg)

					writeRegistry(i, pkg)

					var pkgToSend *ProtoLogistic.Package
					pkgToSend = pkg.pkg
					pkgToSend.Estado = pkg.estado
					pkgToSend.Intentos = int32(pkg.tries)
					_, err = truckService.FinishPackage(context.Background(), pkgToSend)
					if err != nil {
						panic(err)
					}
				}

				truck.pkgsDone = nil
				truck.pkgs = nil
				truck.pkgsToDeliver = nil
			}
		}(i)
	}
	<-forever
	fmt.Println("Tiempo de cerrar.")
}

// Truck : Estructura de estados de un camion
type Truck struct {
	Type          ProtoLogistic.TruckType
	pkgs          []*ProtoLogistic.Package
	pkgsToDeliver []*pkgOnDeliver
	pkgsDone      []*pkgOnDeliver
	registry      []*pkgOnDeliver
}

type pkgOnDeliver struct {
	pkg      *ProtoLogistic.Package
	fecha    string
	estado   string
	maxTries int
	tries    int
}

func (pkg *pkgOnDeliver) addATry() {
	pkg.tries++
}

func (pkg *pkgOnDeliver) checkMaxTries() bool {
	return pkg.maxTries == pkg.tries
}

func (pkg *pkgOnDeliver) setStatus(status string) {
	pkg.estado = status
}

func (pkg *pkgOnDeliver) setDeliveredDate(right bool) {
	if right {
		pkg.fecha = time.Now().Format("2006.01.02 15:04:05")
	} else {
		pkg.fecha = "0"
	}
}

func writeRegistry(i int, pkg *pkgOnDeliver) {
	filename := "registroCamion" + strconv.Itoa(i) + ".csv"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		file, err = os.Create(filename)
		if err != nil {
			panic(err)
		}
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escritura en Registro
	var tipo string
	if pkg.pkg.GetTipo() == 1 {
		tipo = "retail"
	} else if pkg.pkg.GetTipo() == 2 {
		tipo = "prioritario"
	} else {
		tipo = "normal"
	}
	toWrite := []string{
		pkg.pkg.GetIDPaquete(),
		tipo,
		strconv.Itoa(int(pkg.pkg.GetValor())),
		pkg.pkg.GetOrigen(),
		pkg.pkg.GetDestino(),
		strconv.Itoa(pkg.tries),
		pkg.fecha}
	fmt.Println("Escribirndo: ", toWrite)
	err = writer.Write(toWrite)
	if err != nil {
		fmt.Println(err)
	}
}
