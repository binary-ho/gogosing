package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("GET /{$}", app.homeView)
	serverMux.HandleFunc("GET /body-check", app.responseBodyTestView)
	serverMux.HandleFunc("GET /user/{id}", app.pathParameterView)
	serverMux.HandleFunc("GET /json", app.jsonView)

	return serverMux
}
