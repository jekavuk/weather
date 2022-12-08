package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"weatherclient"

	"github.com/joho/godotenv"
)

func main() {

	args := os.Args

	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalf("loading .env: %s", err)
	}

	location := args[1]
	key := os.Getenv("API_KEY")
	url := weatherclient.FormatURL(location, key)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("unexpected response status", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	conditions, err := weatherclient.ParseResponse(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", conditions)
}
