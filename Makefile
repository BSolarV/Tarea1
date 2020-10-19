all: Logistica Clientes Camiones Finanzas

Logistica: Logistica/Logistica.go
	go build -o ./Logistica ./Logistica/Logistica.go

Clientes: 
	go build -o ./Clientes ./Clientes/Clientes.go

Camiones: 
	go build -o ./Camiones ./Camiones/Camiones.go

Finanzas: 
	go build -o ./Finanzas ./Finanzas/Finanzas.go