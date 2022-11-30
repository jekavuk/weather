package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"weatherclient"
)

func main() {

	l := "London"
	key := "dummyKey"
	url := weatherclient.FormatURL(l, key)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("unexpected response status", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	conditions, err := weatherclient.ParseResponse(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conditions)
}
