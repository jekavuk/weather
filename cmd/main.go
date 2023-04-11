package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/jekvuk/weather"
)

// options
var temperatureScale = flag.String("scale", "", "port number")

func main() {
	flag.Parse()

	locationArgs := flag.Args()
	if len(locationArgs) < 1 {
		log.Fatal("please provide valid location")
	}

	key, err := weather.GetAPIKey()
	if err != nil {
		log.Fatal(err)
	}

	weClient := weather.NewClient(key)
	weClient.TemperatureScale = *temperatureScale

	location := strings.Join(locationArgs, " ")
	conditions, err := weClient.GetWeather(location)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", conditions)
}
