all: Logistica Clientes Camiones Finanzas

UpdateGo: 
	export GOROOT=/usr/local/go 
	export GOPATH=$HOME/Projects/Proj1 
	export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

BuildLogistica: Logistica/Logistica.go
	make UpdateGo
ifneq ("$(wildcard $(./bin/Logistica))","")
	rm -r ./bin
endif
	go build -o ./bin/Logistica ./Logistica/Logistica.go

BuildClientes: Clientes/Clientes.go
	make UpdateGo
ifneq ("$(wildcard $(bin/Clientes))","")
	rm -r ./bin
endif
	go build -o ./bin/Clientes ./Clientes/Clientes.go

BuildCamiones: Camiones/Camiones.go
	make UpdateGo
ifneq ("$(wildcard $(bin/Camiones))","")
	rm -r ./bin
endif
	go build -o ./bin/Camiones ./Camiones/Camiones.go

BuildFinanzas: Finanzas/Finanzas.go
	make UpdateGo
ifneq ("$(wildcard $(bin/Finanzas))","")
	rm -r ./bin
endif
	go build -o ./bin/Finanzas ./Finanzas/Finanzas.go

clearRegisters:
	rm *.csv