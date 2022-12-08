package weatherclient

import (
	"encoding/json"
	"fmt"
)

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

func ParseResponse(data []byte) (APIResponse, error) {
	var apiResp APIResponse
	err := json.Unmarshal(data, &apiResp)
	if err != nil {
		return APIResponse{}, err
	}
	return apiResp, nil
}
