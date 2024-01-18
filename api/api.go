package api

import (
    "fmt"
    "net/http"
	"log"
	"io"
)

func Get(apiKey string, start_date string, end_date string, ticker string)  string{
	var url string = buildURL(apiKey, start_date, end_date, ticker)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))

	return string(body)
}

func buildURL(apiKey string, start_date string, end_date string, ticker string) string {
	var BuildedURL string = fmt.Sprintf("https://api.polygon.io/v2/aggs/ticker/%s/range/1/day/%s/%s?apiKey=%s", ticker, start_date, end_date, apiKey)

	return BuildedURL
}