package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const APIKeyName = "OPEN_WEATHER_MAP_API_KEY"

type WeatherClient struct {
	APIKey, BaseURL, TemperatureScale string
	HTTPClient                        *http.Client
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

func (wc WeatherClient) GetWeather(location string) (string, error) {
	URL := wc.FormatURL(location)
	resp, err := wc.HTTPClient.Get(URL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	conditions, err := ParseResponse(data, wc.TemperatureScale)
	if err != nil {
		return "", err
	}

	return conditions, nil
}

func ParseResponse(data []byte, temperatureScale string) (string, error) {
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
		return "", err
	}
	if len(apiResp.Weather) < 1 {
		return "", fmt.Errorf("invalid weather data: %q", data)
	}
	var tempC float64
	var tempSymbol string
	switch temperatureScale {
	case "celsius":
		tempC = apiResp.Main.Temp - 273.5
		tempSymbol = "ºC"
	case "fahrenheit":
		tempC = 1.8*(apiResp.Main.Temp-273.15) + 32
		tempSymbol = "ºF"
	default:
		tempC = apiResp.Main.Temp
		tempSymbol = "K"
	}
	return fmt.Sprintf("Current wether for %s: %s %.1f%s", apiResp.Name, apiResp.Weather[0].Main, tempC, tempSymbol), nil
}

func GetAPIKey() (string, error) {
	key := os.Getenv(APIKeyName)
	if key == "" {
		return "", fmt.Errorf("please set env var %s to a value of your API key", APIKeyName)
	}
	return key, nil
}
