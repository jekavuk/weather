# Description
`weather` is a simple Go library package with accompanying command-line tool. It provides a short information about current weather conditions for a given location using the [OpenWeatherMap API](https://openweathermap.org/api).

# Installation
If you wish install it using go get:
```
go get github.com/jekavuk/weather
```

# Usage
To use weather libraty, import it:
```
import "github.com/jekavuk/weather"
```
Then, you can create a new client and use its GetWeather method to retrieve weather information for a specific location:
```
client := weatherclient.NewClient("API_KEY")
weather, err := client.GetWeather("New York")
if err != nil {
    // handle error
}
// use weather information
```
You will need to replace API_KEY with your own API key from OpenWeatherMap. You can sign up for an API key for free at https://openweathermap.org/api.

## Frstures
There is a feature option to display current temperature as Celsius, Fahrenheit or default(Kelvin).
To try it out provide option parameter *scale* when running the tool for example provide option like:
```
--scale=celsius
```
Possible values are lower case strings (celsius or fahrenheit). If this option is not set, temperature will be displayed for default temperature scale.
If you are using library package set this option on weather client (in the same maner).

# License
This project is licensed under the MIT License - see the LICENSE.md file for details.
