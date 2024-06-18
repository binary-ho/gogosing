package main

import (
	"fmt"
	"net/http"
	"strconv"
)

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
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("GET /{$}", homeView)
	serverMux.HandleFunc("GET /body-check", responseBodyTestView)
	serverMux.HandleFunc("GET /user/{id}", pathParameterView)
	serverMux.HandleFunc("GET /json", jsonView)

	fmt.Println("Hello Guest Book! Starting Server On -> ", PortNumber)

	err := http.ListenAndServe(getAddress(), serverMux)
	fmt.Println("error : ", err)
}

func getAddress() string {
	return ":" + strconv.Itoa(PortNumber)
}
