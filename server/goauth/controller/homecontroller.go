package controller

import (
	"gogosing/server/goauth/ui"
	"net/http"
)

type HomeController struct{}

func (controller *HomeController) Serve(mux *http.ServeMux) {
	mux.HandleFunc(controller.renderMain())
}

const HomeEndpoint = "/{$}"

func (controller *HomeController) renderMain() (string, func(http.ResponseWriter, *http.Request)) {
	return HomeEndpoint, func(writer http.ResponseWriter, request *http.Request) {
		ui.Render(writer, "main", nil)
	}
}
