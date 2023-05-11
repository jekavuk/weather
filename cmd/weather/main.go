package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/jekavuk/weather"
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

	location := strings.Join(locationArgs, " ")
	conditions, err := weClient.GetWeather(location)
	if err != nil {
		log.Fatal(err)
	}

	switch *temperatureScale {
	default:
		fmt.Println(conditions.StringCelsius())
	case "fahrenheit":
		fmt.Println(conditions.StringFahrenheit())
	}
}
