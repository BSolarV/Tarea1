all: MainLogistica MainClientes MainCamiones MainFinanzas

MainLogistica: Logistica/Logistica.go
	go build -o ./bin/MainLogistica ./Logistica/Logistica.go

MainClientes: Logistica/Logistica.go
	go build -o ./bin/MainClientes ./Clientes/Clientes.go

MainCamiones: Camiones/Camiones.go
	go build -o ./bin/MainCamiones ./Camiones/Camiones.go

MainFinanzas: Finanzas/Finanzas.go
	go build -o ./bin/MainFinanzas ./Finanzas/Finanzas.go