package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jekvuk/weather"
)

// options
var temperatureScale = flag.String("scale", "", "port number")

func main() {
	var locationArgs []string
	if os.Args != nil && len(os.Args) > 2 {
		locationArgs = append(locationArgs, os.Args[2:]...)
	}
	if len(locationArgs) < 1 {
		log.Fatal("please provide valid location")
	}

	key, err := weather.GetAPIKey()
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()

	weClient := weather.NewClient(key)
	weClient.TemperatureScale = *temperatureScale

	location := strings.Join(locationArgs, " ")
	conditions, err := weClient.GetWeather(location)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", conditions)
}
