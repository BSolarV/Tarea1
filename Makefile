all: Logistica Clientes Camiones Finanzas

Logistica: Logistica/Logistica.go clean
	go build -o ./bin/Logistica ./Logistica/Logistica.go

Clientes: Logistica/Logistica.go clean
	go build -o ./bin/Clientes ./Clientes/Clientes.go

Camiones: Camiones/Camiones.go clean
	go build -o ./bin/Camiones ./Camiones/Camiones.go

Finanzas: Finanzas/Finanzas.go clean
	go build -o ./bin/Finanzas ./Finanzas/Finanzas.go

clean:
	rm -r ./bin