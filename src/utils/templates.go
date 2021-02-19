package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// LoadTemplates parses all HTML templates and add them into the templates variable
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// RenderTemplate renders an HTML page on the client's browser
func RenderTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}