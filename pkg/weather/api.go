package weather

import (
	"net/http"
)

// Client struct represents the client connecting to the Weather API
type Client struct {
	http *http.Client
	key  string
}

// NewClient creates and returns a new Client instance for making requests to the Weather API
func NewClient(httpClient *http.Client, key string) *Client {
	return &Client{httpClient, key}
}