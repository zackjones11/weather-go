package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/zackjones11/weather-go/pkg/photo"
	"github.com/zackjones11/weather-go/pkg/weather"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error: Cannot load .env file")
	}

	client := &http.Client{Timeout: 10 * time.Second}

	weatherAPIKey := os.Getenv("WEATHER_API_KEY")
	if weatherAPIKey == "" {
		log.Fatal("Error: Please add a WEATHER_API_KEY to .env")
	}

	weatherAPI := weather.NewClient(client, weatherAPIKey)

	photoAPIKey := os.Getenv("PHOTO_API_KEY")
	if photoAPIKey == "" {
		log.Fatal("Error: Please add a PHOTO_API_KEY to .env")
	}

	photoAPI := photo.NewClient(client, photoAPIKey)

	assetsFs := http.FileServer(http.Dir("public/assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", assetsFs))
	mux.HandleFunc("/", weather.SearchHandler)
	mux.HandleFunc("/weather", weather.DetailHandler(photoAPI, weatherAPI))
	http.ListenAndServe(":8080", mux)
}
