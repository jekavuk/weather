package weatherclient_test

import (
	"fmt"
	"os"
	"testing"
	wc "weatherclient"

	"github.com/google/go-cmp/cmp"
)

func TestFormatURL_ReturnsCorrectURLForProvidedLocationAndAPIKey(t *testing.T) {

	want := "https://api.openweathermap.org/data/2.5/weather?q=Belgrade&appid=dummyKey"
	got := wc.FormatURL("Belgrade", "dummyKey")

	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseResponse_CorrectlyParsesResponseIntoString(t *testing.T) {
	want := "Clear 6.8ºC"
	var data []byte
	data, err := os.ReadFile("testdata/bgwether.json")
	if err != nil {
		t.Fatal(err)
	}

	got, err := wc.ParseResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseResponse_ReturnsErrorForInvalidJSON(t *testing.T) {
	_, err := wc.ParseResponse([]byte{})
	if err == nil {
		t.Fatal("wanted error for invalid data and got nil")
	}
}

func TestGetAPIKey_CorrectlyGetsAPIKeyIfEnvVarIsSet(t *testing.T) {
	want := "dummyKey"
	t.Setenv(wc.APIKeyName, want)
	got, err := wc.GetAPIKey()
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestGETAPIKey_ReturnsErrorForMissingEnvVar(t *testing.T) {
	t.Setenv(wc.APIKeyName, "")
	_, err := wc.GetAPIKey()
	if err == nil {
		t.Fatal("wanted error for missing env var")
	}
}

func TestGetWeather_CorrectlyReturnsWeatherInfo(t *testing.T) {
	want := "Clear 6.8ºC"
	weClient := wc.NewWeatherClient("dummyKey", "Belgrade")
	got, err := weClient.GetWeather()
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestNewWeatherClient_ReturnsInstanceOfWeatherClientWithSetURLAndAPIKey(t *testing.T) {
	location := "Belgrade"
	apiKey := "dummyKey"

	want := wc.WeatherClient{
		URL: fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", location, apiKey),
	}
	got := wc.NewWeatherClient(apiKey, location)

	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}
