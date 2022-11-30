package weatherclient_test

import (
	"os"
	"testing"
	"weatherclient"

	"github.com/google/go-cmp/cmp"
)

func TestFormatURL(t *testing.T) {

	want := "https://api.openweathermap.org/data/2.5/weather?q=Belgrade&appid=dummyKey"
	got := weatherclient.FormatURL("Belgrade", "dummyKey")

	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseResponse(t *testing.T) {
	want := "Clear"
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

func TestParseInvalidResponse(t *testing.T) {
	_, err := weatherclient.ParseResponse([]byte{})
	if err == nil {
		t.Fatal("wanted error for invalid data and got nil")
	}
}
