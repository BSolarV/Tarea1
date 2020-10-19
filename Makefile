all: Logistica Clientes Camiones Finanzas

Logistica: bin/Logistica
	make BuildLogistica

Clientes: bin/Clientes
	make BuildClientes

Camiones: bin/Camiones
	make BuildCamiones

Finanzas: bin/Finanzas
	make BuildFinanzas

BuildLogistica: Logistica/Logistica.go
	go build -o ./bin/Logistica ./Logistica/Logistica.go

BuildClientes: Clientes/Clientes.go
	go build -o ./bin/Clientes ./Clientes/Clientes.go

BuildCamiones: Camiones/Camiones.go
	go build -o ./bin/Camiones ./Camiones/Camiones.go

BuildFinanzas: Finanzas/Finanzas.go
	go build -o ./bin/Finanzas ./Finanzas/Finanzas.go

clean:
	rm -r ./bin

clearRegisters:
	rm *.csv