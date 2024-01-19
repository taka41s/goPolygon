package api

import(
  "fmt"
  "capybara/config"
  "log"
)

func Aggs(ticker string, start_date string, end_date string) string {
	env, err := config.LoadDotenv()

	if err != nil {
		log.Fatalln(err)
	}

	var BuildedURL string = fmt.Sprintf("%s/aggs/ticker/%s/range/1/day/%s/%s?apiKey=%s", env.POLYGON_BASEURL, ticker, start_date, end_date, env.POLYGON_APIKEY)

	return BuildedURL
}