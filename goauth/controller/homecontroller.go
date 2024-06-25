package controller

import "net/http"

type HomeController struct{}

func (controller HomeController) Serve(mux *http.ServeMux) {
	mux.HandleFunc(renderMain())
	mux.HandleFunc(renderAuth())
	mux.HandleFunc(authentication())
}

func renderMain() (string, func(http.ResponseWriter, *http.Request)) {
	return "/", func(w http.ResponseWriter, r *http.Request) {}
}

func renderAuth() (string, func(http.ResponseWriter, *http.Request)) {
	return "/oauth", func(w http.ResponseWriter, r *http.Request) {}
}

func authentication() (string, func(http.ResponseWriter, *http.Request)) {
	return "/oauth/callback", func(w http.ResponseWriter, r *http.Request) {}
}
