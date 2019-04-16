package api

import (
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
	req = setHeaders(req)

	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

// setHeaders is a workaround because there are two
// different api's. TODO: ughh refactor this.
func setHeaders(req *http.Request) *http.Request {
	req.Header.Set("User-Agent", UserAgent)
	if req.URL.Host == "data.net" {
		req.Header.Set("Referer", "http://data.nba.net/")
		req.Header.Set("Origin", "http://data.nba.net")
	} else {
		req.Header.Set("Referer", Referer)
		req.Header.Set("Origin", Origin)
	}
	return req
}
