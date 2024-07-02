package ui

import (
	"gogosing/server/goauth/util"
	"html/template"
	"net/http"
)

func Render(writer http.ResponseWriter, htmlFileName string, data interface{}) {
	tmpl, _ := template.ParseFiles(util.ProjectPath + "ui\\" + htmlFileName + ".html")
	tmpl.Execute(writer, data)
}
