package main

import (
	"log/slog"
	"net/http"
	"os"
)

type Application struct {
	logger *slog.Logger
}

var App *Application

func init() {
	App = &Application{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

func (app *Application) Run() {
	err := http.ListenAndServe("localhost:8080", app.routes())
	if err != nil {
		app.logger.Error(err.Error())
	}
}
