package weatherclient

import (
	"encoding/json"
	"fmt"
)

type APIResponse struct {
	Weather []struct {
		Main string
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
	return apiResp.Weather[0].Main, nil
}
