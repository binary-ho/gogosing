package main

import (
	"github.com/gorilla/sessions"
	"gogosing/server/goauth/controller"
	"log/slog"
	"net/http"
	"os"
)

type Application struct {
	logger       *slog.Logger
	sessionStore *sessions.CookieStore
	servers      *[]Server
}

var App *Application

func init() {
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

	App = &Application{
		logger:       slog.New(slog.NewTextHandler(os.Stdout, nil)),
		sessionStore: store,
		servers: &[]Server{
			&controller.HomeController{},
			&controller.AuthController{
				SessionStore: store,
			},
		},
	}
}

func (app *Application) Run() {
	err := http.ListenAndServe("localhost:8080", app.routes())
	if err != nil {
		app.logger.Error(err.Error())
	}
}
