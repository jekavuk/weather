package weatherclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const APIKeyName = "OPEN_WEATHER_MAP_API_KEY"

type WeatherClient struct {
	URL string
}

type APIResponse struct {
	Weather []struct {
		Main string
	}
	Main struct {
		Temp      float64
		FeelsLike float64
		Pressure  int
		Humidity  int
	}
}

func NewWeatherClient(apiKey, location string) WeatherClient {
	url := FormatURL(location, apiKey)
	return WeatherClient{URL: url}
}

func FormatURL(location string, apiKey string) string {
	return fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", location, apiKey)
}

func ParseResponse(data []byte) (string, error) {
	var apiResp APIResponse
	err := json.Unmarshal(data, &apiResp)
	if err != nil {
		return "", err
	}
	tempC := apiResp.Main.Temp - 273.5
	return fmt.Sprintf("%s %.1fºC", apiResp.Weather[0].Main, tempC), nil
}

func GetAPIKey() (string, error) {
	key := os.Getenv(APIKeyName)
	if key == "" {
		return "", fmt.Errorf("please set env var %s to a value of your API key", APIKeyName)
	}
	return key, nil
}

func (wc WeatherClient) GetWeather() (string, error) {
	resp, err := http.Get(wc.URL)
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

	conditions, err := ParseResponse(data)
	if err != nil {
		return "", err
	}

	return conditions, nil
}
