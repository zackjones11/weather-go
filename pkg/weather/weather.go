package weather

import (
	"html/template"
	"math"
	"net/http"
	"net/url"

	"github.com/zackjones11/weather-go/pkg/photo"
)

// Details contains info to give the template
type Details struct {
	TempActual      float64
	Description     string
	IconName        string
	Location        string
	BackgroundImage string
}

// SearchHandler handles the request to begin searching for weather by location
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	t := loadTemplate(w, "search.html")
	t.Execute(w, nil)
}

// DetailHandler handles the request to display the weather for a specific location
func DetailHandler(photo *photo.Client, weather *Client) http.HandlerFunc {
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

		photoResults, err := photo.GetRandomPhoto(locationQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		weatherDetails := &Details{
			TempActual:      math.Round(results.Main.Temp),
			Description:     results.Weather[0].Description,
			IconName:        results.Weather[0].Main,
			Location:        results.Location,
			BackgroundImage: photoResults.Urls.Regular,
		}

		t.Execute(w, weatherDetails)
	}
}

func loadTemplate(w http.ResponseWriter, path string) *template.Template {
	return template.Must(template.ParseFiles("public/" + path))
}
