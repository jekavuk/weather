package weatherclient_test

import (
	"os"
	"testing"
	"weatherclient"

	"github.com/google/go-cmp/cmp"
)

func TestFormatURL_ReturnsCorrectURLForProvidedLocationAndAPIKey(t *testing.T) {

	want := "https://api.openweathermap.org/data/2.5/weather?q=Belgrade&appid=dummyKey"
	got := weatherclient.FormatURL("Belgrade", "dummyKey")

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

	got, err := weatherclient.ParseResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseResponse_ReturnsErrorForInvalidJSON(t *testing.T) {
	_, err := weatherclient.ParseResponse([]byte{})
	if err == nil {
		t.Fatal("wanted error for invalid data and got nil")
	}
}

func TestGetAPIKey_CorrectlyGetsAPIKeyIfEnvVarIsSet(t *testing.T) {
	want := "dummyKey"
	t.Setenv(weatherclient.APIKeyName, want)
	got, err := weatherclient.GetAPIKey()
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestGETAPIKey_ReturnsErrorForMissingEnvVar(t *testing.T) {
	t.Setenv(weatherclient.APIKeyName, "")
	_, err := weatherclient.GetAPIKey()
	if err == nil {
		t.Fatal("wanted error for missing env var")
	}
}

func TestGetWeather_CorrectlyReturnsWeatherInfo(t *testing.T) {
	want := "Clear 6.8ºC"
	got, err := weatherclient.GetWeather("Belgrade", "dummyKey")
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
