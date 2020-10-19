all: logistica 

logistica: 
	go build -o ./logistica ./Logistica/Logistica.go
