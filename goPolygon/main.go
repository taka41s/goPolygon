package main

import (
	"fmt"
	"capybara/api"
)

func main(){
	url := api.Aggs("AAPL","2023-01-09", "2023-01-09")

	fmt.Println(api.Get(url))
}
