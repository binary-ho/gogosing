package main

import (
	"fmt"
	"net/http"
)

func homeView(writer http.ResponseWriter, _ *http.Request) {
	writeMessage(writer, "Hello! Good afternoon, good evening, and good night.")
}

func responseBodyTestView(writer http.ResponseWriter, _ *http.Request) {
	body := TestResponseBody{field1: "test", field2: 77, innerStruct1: InnerStruct1{"1", 11.11}, innerStruct2: innerStruct2{"2", 22.22}}
	//io.WriteString(writer, body)
	fmt.Fprint(writer, body)
}

func pathParameterView(writer http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")
	writer.WriteHeader(http.StatusOK)
	writeMessage(writer, "입력된 ID : "+id)
}

func jsonView(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writeMessage(writer, `{"name":"Jinho"}`)
}

func writeMessage(w http.ResponseWriter, message string) {
	w.Write([]byte(message))
}
