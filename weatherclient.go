package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const APIKeyName = "OPEN_WEATHER_MAP_API_KEY"

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

func (wc WeatherClient) FormatURL(location string) string {
	return fmt.Sprintf("%s?q=%s&appid=%s", wc.BaseURL, location, wc.APIKey)
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

	conditions, err := ParseResponse(data)
	if err != nil {
		return "", err
	}

	return conditions, nil
}

func ParseResponse(data []byte) (string, error) {
	var apiResp struct {
		Weather []struct {
			Main string
		}
		Main struct {
			Temp float64
		}
	}
	err := json.Unmarshal(data, &apiResp)
	if err != nil {
		return "", err
	}
	if len(apiResp.Weather) < 1 {
		return "", fmt.Errorf("invalid weather data: %q", data)
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
