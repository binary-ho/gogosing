package main

import "net/http"

type Server interface {
	Serve(mux *http.ServeMux)
}
