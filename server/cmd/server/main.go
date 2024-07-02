package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

type application struct {
	logger *slog.Logger
}

type TestResponseBody struct {
	field1       string
	field2       int
	innerStruct1 InnerStruct1
	innerStruct2 innerStruct2
}

type InnerStruct1 struct {
	subField1 string
	subField2 float64
}

type innerStruct2 struct {
	subField1 string
	subField2 float64
}

const PortNumber = 8080

func main() {
	application := &application{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	//addr := os.Getenv
	printServerRun()
	err := http.ListenAndServe(getAddress(), application.routes())
	fmt.Println("error : ", err)
}

func printServerRun() (int, error) {
	return fmt.Println("Hello Guest Book! Starting Server On -> ", PortNumber)
}

func getAddress() string {
	return ":" + strconv.Itoa(PortNumber)
}
