package main

import (
	"fmt"
	"log"
	"os"
	"weatherclient"
)

func main() {
	var location string
	if os.Args != nil && len(os.Args) > 1 {
		location = os.Args[1]
	}
	if location == "" {
		log.Fatal("please provide valid location")
	}

	key, err := weatherclient.GetAPIKey()
	if err != nil {
		log.Fatal(err)
	}

	conditions, err := weatherclient.GetWeather(location, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", conditions)
}
