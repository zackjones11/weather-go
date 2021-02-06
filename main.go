package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/zackjones11/weather-go/pkg/weather"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error: Cannot load .env file")
	}

	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("Error: Please add a WEATHER_API_KEY to .env")
	}

	client := &http.Client{Timeout: 10 * time.Second}
	weatherAPI := weather.NewClient(client, apiKey)

	mux := http.NewServeMux()
	mux.HandleFunc("/", weather.SearchHandler)
	mux.HandleFunc("/weather", weather.DetailHandler(weatherAPI))
	http.ListenAndServe(":8080", mux)
}