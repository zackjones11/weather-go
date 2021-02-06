package weather

import (
	"html/template"
	"net/http"
)

// SearchHandler handles the request to begin searching for weather by location
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("public/index.html"))
	tpl.Execute(w, nil)
}