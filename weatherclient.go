package weatherclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const APIKeyName = "OPEN_WEATHER_MAP_API_KEY"

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

func GetWeather(location string, key string) (string, error) {
	url := FormatURL(location, key)

	resp, err := http.Get(url)
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
