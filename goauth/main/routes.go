package main

import (
	"gogosing/goauth/controller"
	"net/http"
)

func (app *Application) routes() *http.ServeMux {
	serveMux := http.NewServeMux()

	homeController := &controller.HomeController{}
	homeController.Serve(serveMux)
	return serveMux
}
