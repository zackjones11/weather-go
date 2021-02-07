package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client struct represents the client connecting to the Weather API
type Client struct {
	http *http.Client
	key  string
}

// Response contains the whole response from the API
type Response struct {
	Weather  []Weather `json:"weather"`
	Main     Main      `json:"main"`
	Location string    `json:"name"`
}

// Weather contains the information for the weather requested
type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

// Main contains tempatures from the weather requested
type Main struct {
	Temp float64 `json:"temp"`
}

// NewClient creates and returns a new Client instance for making requests to the Weather API
func NewClient(httpClient *http.Client, key string) *Client {
	return &Client{httpClient, key}
}

// GetWeather fetches the current weather for a given location
func (c *Client) GetWeather(location string) (*Response, error) {
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + location + "&appid=" + c.key + "&units=metric"
	response, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	responseObject := &Response{}
	return responseObject, json.Unmarshal(body, responseObject)
}
