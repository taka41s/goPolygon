package config

import (
	"os"
	"github.com/joho/godotenv"
	"log"
)

type Env struct {
	POLYGON_BASEURL string
	POLYGON_APIKEY  string
}

func LoadDotenv() (*Env, error){
	err := godotenv.Load()

	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	env := &Env{
		POLYGON_BASEURL: os.Getenv("POLYGON_BASEURL"),
		POLYGON_APIKEY:  os.Getenv("POLYGON_APIKEY"),
	}
	
	return env, nil
}
