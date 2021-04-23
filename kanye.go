package kanye

import (
	"encoding/json"
	"net/http"
)

// // Kanye wraps the whole kanye.rest api client
// type Kanye struct {
// 	Client Client
// }

// Client wraps a `http.client`
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient initializes a Client with the base api url
// and a default http client
func NewClient() *Client {
	return &Client{
		BaseURL:    "https://api.kanye.rest",
		HTTPClient: http.DefaultClient,
	}
}

// Quote denotes the structure of the response
// which is a json with one field containing the quote
type Quote struct {
	Quote string `json:"quote"`
}

func (c *Client) get(path string, v interface{}) error {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}

// GetQuote returns a random quote
func (c *Client) GetQuote() (string, error) {
	path := c.BaseURL
	var quote Quote

	err := c.get(path, &quote)
	if err != nil {
		return "", err
	}

	return quote.Quote, nil
}
