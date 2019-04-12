package api

import (
	"fmt"
	"net/http"
)

type Client struct {
	NBAService *NBAService
}

type service struct {
	client *Client
}

// NBAService ...
type NBAService service

func NewClient() *Client {
	c := &Client{}
	c.NBAService = &NBAService{}
	return c
}

func apiGet(reqURL string, params map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", reqURL, nil)

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL)
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Referer", Referer)
	req.Header.Set("Origin", Origin)

	if err != nil {
		return nil, err
	}
	return client.Do(req)
}
