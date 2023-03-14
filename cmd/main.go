package main

import (
	"fmt"
	"log"
	"os"
	wc "weatherclient"
)

func main() {
	var location string
	if os.Args != nil && len(os.Args) > 1 {
		location = os.Args[1]
	}
	if location == "" {
		log.Fatal("please provide valid location")
	}

	key, err := wc.GetAPIKey()
	if err != nil {
		log.Fatal(err)
	}

	weClient := wc.NewWeatherClient(key, location)

	conditions, err := weClient.GetWeather()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", conditions)
}
