package main

import (
	"net/http"
)

func (app *Application) routes() *http.ServeMux {
	serveMux := http.NewServeMux()
	for _, server := range *app.servers {
		server.Serve(serveMux)
	}
	return serveMux
}
