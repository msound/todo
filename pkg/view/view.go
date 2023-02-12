package view

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func Load() {
	templates = template.Must(template.ParseFiles("index.html"))
}

func Render(w http.ResponseWriter, template string, data any) {
	templates.ExecuteTemplate(w, template, data)
}
