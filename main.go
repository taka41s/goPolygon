package main

import (
	// "errors"
	"fmt"
	"capybara/api"
	// "net/http"
	// "os"
	// "time"
)

func main(){
	var apiKey string = "ogKoJqWTjDIkOTLdVpPsJv9pbkwjNuOv"
	fmt.Println(api.Get(apiKey, "2023-01-09", "2023-01-09", "AAPL"))
}