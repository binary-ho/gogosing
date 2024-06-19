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
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("GET /{$}", application.homeView)
	serverMux.HandleFunc("GET /body-check", application.responseBodyTestView)
	serverMux.HandleFunc("GET /user/{id}", application.pathParameterView)
	serverMux.HandleFunc("GET /json", application.jsonView)

	fmt.Println("Hello Guest Book! Starting Server On -> ", PortNumber)

	err := http.ListenAndServe(getAddress(), serverMux)
	fmt.Println("error : ", err)
}

func getAddress() string {
	return ":" + strconv.Itoa(PortNumber)
}
