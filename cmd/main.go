package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"weatherclient"

	"github.com/joho/godotenv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	err := godotenv.Load("config.env")
	if err != nil {
		log.Printf("config file missing: %s", err)
		// API key
		for {
			fmt.Print("Please enter youre API key: ")
			scanner.Scan()
			text := scanner.Text()

			if len(text) == 0 {
				fmt.Println("API key must be provided")
				continue
			} else {
				os.Setenv("API_KEY", text)
				break
			}
		}
	}

	var location string
	if os.Args != nil && len(os.Args) > 1 {
		location = os.Args[1]
	}

	if location == "" {
		// location
		for {
			fmt.Print("Please enter location: ")
			scanner.Scan()
			text := scanner.Text()

			if len(text) == 0 {
				fmt.Println("Location must be provided")
				continue
			} else {
				location = text
				break
			}
		}
	}

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
