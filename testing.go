package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("result.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		file, err = os.Create("result.csv")
		if err != nil {
			panic(err)
		}
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	end := make(chan bool)
	go func() {
		var toWrite []string
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Ingrese valor -> ")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			text = strings.Replace(text, "\r", "", -1)
			if text == "" {
				err := writer.Write(toWrite)
				checkError("Cannot write to file", err)
				toWrite = nil
			} else if text == "FIN" {
				end <- true
			} else {
				toWrite = append(toWrite, text)
			}
		}
	}()
	<-end
}

func checkError(message string, err error) {
	if err != nil {
		fmt.Println(message, err)
	}
}
