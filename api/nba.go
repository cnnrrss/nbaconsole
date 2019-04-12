package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	siteURL        = "https://stats.nba.com"
	edge           = "stats"
	endpoint       = "scoreboardV2"
	gameDateFormat = "2006-01-02"
)

type Client struct {
	NBAService *NBAService
}

type service struct {
	client *Client
}

// NBAService ...
type NBAService service

// Response is the response structure
type Response struct {
	// Status Status      `json:"status"`
	ResultSets interface{} `json:"resultSets"`
}

func NewClient() *Client {
	c := &Client{}
	c.NBAService = &NBAService{}
	return c
}

func FetchNBAToday() (string, error) {
	client := http.Client{Timeout: 10 * time.Second}
	root := "http://data.nba.net/10s/prod/v1/today.json"
	res, err := client.Get(root)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	res.Body.Close()
	return string(body), nil
}

func (c *Client) TodaysGames(options string) (string, error) {
	date := time.Now().AddDate(0, 0, -1).Format(gameDateFormat)
	// https://stats.nba.com/stats/scoreboardV2?DayOffset=0&LeagueID=00&gameDate=05%2F01%2F2019
	url := fmt.Sprintf("%s/%s/%s?DayOffset=0&LeagueID=00&gameDate=%s", siteURL, edge, endpoint, date)

	body, err := makeReq(url)
	if err != nil {
		return "", fmt.Errorf("Error with request %v", err)
	}

	//TODO: unmarshal responsebody

	return string(body), nil
}

// makeReq HTTP request helper
func makeReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// TODO: refactor headers
	req.Header.Add("Host", "stats.nba.com")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "http://stats.nba.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Add("cache-control", "no-cache")

	respBody, err := doReq(req)
	if err != nil {
		return nil, err
	}

	return respBody, err
}

// doReq HTTP client
func doReq(req *http.Request) ([]byte, error) {
	requestTimeout := time.Duration(10 * time.Second)
	tr := &http.Transport{
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   requestTimeout,
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// TODO: refactor defer
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}
