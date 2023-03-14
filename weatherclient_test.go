package weather_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jekvuk/weather"

	"github.com/google/go-cmp/cmp"
)

func TestFormatURL_ReturnsCorrectURLForProvidedLocationAndAPIKey(t *testing.T) {
	wc := weather.NewClient("dummyKey")
	want := "https://api.openweathermap.org/data/2.5/weather?q=Belgrade&appid=dummyKey"
	got := wc.FormatURL("Belgrade")

	if want != got {
		t.Errorf("\nwant %q\ngot  %q", want, got)
	}
}

func TestFormatURL_EscapesSpacesInLocation(t *testing.T) {
	wc := weather.NewClient("dummyKey")
	want := "https://api.openweathermap.org/data/2.5/weather?q=New%20York%20City&appid=dummyKey"
	got := wc.FormatURL("New York City")
	if want != got {
		t.Errorf("\nwant %q\ngot  %q", want, got)
	}
}

func TestFormatURL_HonoursBaseURLSetting(t *testing.T) {
	wc := weather.NewClient("dummyKey")
	wc.BaseURL = "https://example.com/bogusAPI"
	want := "https://example.com/bogusAPI?q=Belgrade&appid=dummyKey"
	got := wc.FormatURL("Belgrade")
	if want != got {
		t.Errorf("\nwant %q\ngot  %q", want, got)
	}
}

func TestParseResponse_CorrectlyParsesResponseIntoString(t *testing.T) {
	want := "Clear 6.8ºC"
	var data []byte
	data, err := os.ReadFile("testdata/bgwether.json")
	if err != nil {
		t.Fatal(err)
	}

	got, err := weather.ParseResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseResponse_ReturnsErrorForInvalidJSON(t *testing.T) {
	_, err := weather.ParseResponse([]byte{})
	if err == nil {
		t.Fatal("wanted error for invalid data and got nil")
	}
}

func TestParseResponse_ReturnsErrorForValidJSONExpressingInvalidData(t *testing.T) {
	_, err := weather.ParseResponse([]byte(`{"bogus":"data"}`))
	if err == nil {
		t.Fatal("wanted error for invalid weather data and got nil")
	}
}

func TestGetAPIKey_CorrectlyGetsAPIKeyIfEnvVarIsSet(t *testing.T) {
	want := "dummyKey"
	t.Setenv(weather.APIKeyName, want)
	got, err := weather.GetAPIKey()
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestGetAPIKey_ReturnsErrorForMissingEnvVar(t *testing.T) {
	t.Setenv(weather.APIKeyName, "")
	_, err := weather.GetAPIKey()
	if err == nil {
		t.Fatal("wanted error for missing env var")
	}
}

func TestGetWeather_CorrectlyReturnsWeatherInfo(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("testdata/bgwether.json")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()
	fmt.Println("server", ts.URL)
	client := weather.NewClient("dummyKey")
	client.BaseURL = ts.URL
	client.HTTPClient = ts.Client()
	want := "Clear 6.8ºC"
	got, err := client.GetWeather("dummyLocation")
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
