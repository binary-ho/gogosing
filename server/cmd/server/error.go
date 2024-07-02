package main

import (
	"net/http"
	"runtime/debug"
)

func (app *application) internal(writer http.ResponseWriter, request *http.Request, err error) {
	var (
		method = request.Method
		uri    = request.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	app.logger.Error(err.Error(), ", method : ", method, ", uri : ", uri, ", trace : ", trace)
	internalError := http.StatusInternalServerError
	http.Error(writer, http.StatusText(internalError), internalError)
}

func (app *application) client(writer http.ResponseWriter, status int) {
	http.Error(writer, http.StatusText(status), status)
}
