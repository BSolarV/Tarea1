all: Logistica 

Logistica: ./Logistica/Logistica.go
	go build -o ./Logistica ./Logistica/Logistica.go
