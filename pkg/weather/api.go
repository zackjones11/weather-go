package weather

import (
	"net/http"
)

// Client struct represents the client connecting to the Weather API
type Client struct {
	http *http.Client
	key  string
}

// Response contains the whole response from the API
type Response struct {
	Weather []Weather `json:"weather"`
}

// Weather contains the information for the weather requested
type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

// NewClient creates and returns a new Client instance for making requests to the Weather API
func NewClient(httpClient *http.Client, key string) *Client {
	return &Client{httpClient, key}
}