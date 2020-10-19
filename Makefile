all: Logistica Clientes Camiones Finanzas

Logistica: Logistica
	make BuildLogistica

Clientes: Clientes
	make BuildClientes

Camiones: Camiones
	make BuildCamiones

Finanzas: Finanzas
	make BuildFinanzas

BuildLogistica: Logistica/Logistica.go
	go build -o ./bin/Logistica ./Logistica/Logistica.go

BuildClientes: Clientes/Clientes.go
	go build -o ./bin/Clientes ./Clientes/Clientes.go

BuildCamiones: Camiones/Camiones.go
	go build -o ./bin/Camiones ./Camiones/Camiones.go

BuildFinanzas: Finanzas/Finanzas.go
	go build -o ./bin/Finanzas ./Finanzas/Finanzas.go

Buildclean:
	rm -r ./bin

BuildclearRegisters:
	rm *.csv