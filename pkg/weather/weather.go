package weather

import (
	"html/template"
	"net/http"
	"net/url"
)

// SearchHandler handles the request to begin searching for weather by location
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	t := loadTemplate(w, "index.html")
	t.Execute(w, nil)
}

// DetailHandler handles the request to display the weather for a specific location
func DetailHandler(weather *Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := loadTemplate(w, "detail.html")

		url, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		params := url.Query()
		locationQuery := params.Get("l")
		results, err := weather.GetWeather(locationQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Execute(w, results.Weather[0])
	}
}

func loadTemplate(w http.ResponseWriter, path string) *template.Template {
	return template.Must(template.ParseFiles("public/" + path))
}