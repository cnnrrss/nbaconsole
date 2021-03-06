package api

import (
	"net/http"
)

func apiGet(reqURL string, params map[string]string) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()
	req = setHeaders(req)

	return client.Do(req)
}

// setHeaders is a workaround because there are two
// different api's. TODO: ughh refactor this.
func setHeaders(req *http.Request) *http.Request {
	req.Header.Set("User-Agent", _userAgent)
	if req.URL.Host == "data.net" {
		req.Header.Set("Referer", "http://data.nba.net/")
		req.Header.Set("Origin", "http://data.nba.net")
	} else {
		req.Header.Set("Referer", _referrer)
		req.Header.Set("Origin", _origin)
	}
	return req
}
