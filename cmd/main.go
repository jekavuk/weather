package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jekvuk/weather"
)

func main() {
	var location string
	if os.Args != nil && len(os.Args) > 1 {
		location = os.Args[1]
	}
	if location == "" {
		log.Fatal("please provide valid location")
	}

	key, err := weather.GetAPIKey()
	if err != nil {
		log.Fatal(err)
	}

	weClient := weather.NewClient(key)

	conditions, err := weClient.GetWeather(location)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", conditions)
}
