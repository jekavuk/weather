package weather

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const APIKeyName = "OPEN_WEATHER_MAP_API_KEY"

type Conditions struct {
	City  string
	Feel  string
	TempK float64
}

type WeatherClient struct {
	APIKey, BaseURL string
	HTTPClient      *http.Client
}

func NewClient(apiKey string) WeatherClient {
	return WeatherClient{
		APIKey:     apiKey,
		BaseURL:    "https://api.openweathermap.org/data/2.5/weather",
		HTTPClient: http.DefaultClient,
	}
}

func (wc WeatherClient) FormatLocation(location string) string {
	return url.PathEscape(location)
}

func (wc WeatherClient) FormatURL(location string) string {
	l := wc.FormatLocation(location)
	return fmt.Sprintf("%s?q=%s&appid=%s", wc.BaseURL, l, wc.APIKey)
}

func (wc WeatherClient) GetWeather(location string) (Conditions, error) {
	URL := wc.FormatURL(location)
	resp, err := wc.HTTPClient.Get(URL)
	if err != nil {
		return Conditions{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Conditions{}, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Conditions{}, err
	}

	conditions, err := ParseResponse(data)
	if err != nil {
		return Conditions{}, err
	}

	return conditions, nil
}

func ParseResponse(data []byte) (Conditions, error) {
	var apiResp struct {
		Weather []struct {
			Main string
		}
		Main struct {
			Temp float64
		}
		Name string
	}
	err := json.Unmarshal(data, &apiResp)
	if err != nil {
		return Conditions{}, err
	}
	if len(apiResp.Weather) < 1 {
		return Conditions{}, fmt.Errorf("invalid weather data: %q", data)
	}

	resp := Conditions{City: apiResp.Name, Feel: apiResp.Weather[0].Main, TempK: apiResp.Main.Temp}
	return resp, nil
}

func GetAPIKey() (string, error) {
	key := os.Getenv(APIKeyName)
	if key == "" {
		return "", fmt.Errorf("please set env var %s to a value of your API key", APIKeyName)
	}
	return key, nil
}

func Main() int {
	var temperatureScale = flag.String("scale", "", "port number")
	flag.Parse()

	locationArgs := flag.Args()
	if len(locationArgs) < 1 {
		fmt.Fprintln(os.Stderr, "please provide valid location")
		return 1
	}

	key, err := GetAPIKey()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	weClient := NewClient(key)

	location := strings.Join(locationArgs, " ")
	conditions, err := weClient.GetWeather(location)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	switch *temperatureScale {
	default:
		fmt.Println(conditions.StringCelsius())
	case "fahrenheit":
		fmt.Println(conditions.StringFahrenheit())
	}

	return 0
}

func (c Conditions) String() string {
	return fmt.Sprintf("Current wether for %s: %s %.1fK", c.City, c.Feel, c.TempK)
}

func (c Conditions) TempFahrenheit() float64 {
	return 1.8*c.TempCelsius() + 32
}

func (c Conditions) StringFahrenheit() string {
	return fmt.Sprintf("Current wether for %s: %s %.1fºF", c.City, c.Feel, c.TempFahrenheit())
}

func (c Conditions) TempCelsius() float64 {
	return c.TempK - 273.5
}

func (c Conditions) StringCelsius() string {
	return fmt.Sprintf("Current wether for %s: %s %.1fºC", c.City, c.Feel, c.TempCelsius())
}
