package photo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client struct represents the client connecting to the Photo API
type Client struct {
	http *http.Client
	key  string
}

// Response contains the whole response from the API
type Response struct {
	Urls Urls `json:"urls"`
}

// Urls contains CDN urls for photos
type Urls struct {
	Regular string `json:"regular"`
}

// NewClient creates and returns a new Client instance for making requests to the Photo API
func NewClient(httpClient *http.Client, key string) *Client {
	return &Client{httpClient, key}
}

// GetRandomPhoto fetches a random photo with a given search query
func (c *Client) GetRandomPhoto(query string) (*Response, error) {
	url := "https://api.unsplash.com/photos/random?query=" + query + "&orientation=landscape" + "&content_filter=high" + "&client_id=" + c.key
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
