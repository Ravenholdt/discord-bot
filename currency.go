package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const apiKey = "9gzE3BpVlwmKXxUC0FtUWNQxqxG3Y6Ib"

type ConversionResult struct {
	Rates map[string]float64 `json:"rates"`
}

var lastUpdate time.Time

func updateRates(base, symbols string) (map[string]float64, error) {
	url := fmt.Sprintf("https://api.apilayer.com/fixer/latest?base=%s&symbols=%s", base, symbols)
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error while creating request: %v", err)
	}

	request.Header.Add("apikey", apiKey)

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error while fetching exchange rates: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading response body: %v", err)
	}

	var result ConversionResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling response: %v", err)
	}

	return result.Rates, nil
}

func currecy() {
	base := "EUR"
	symbols := "SEK,DKK,CZK,GBP"

	result, err := updateRates(base, symbols)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func runAfterTwentyHours() {
	for {
		now := time.Now()
		twentyHoursLater := now.Add(20 * time.Hour)
		if now.Before(twentyHoursLater) {
			fmt.Println("Waiting for 20 hours...")
			time.Sleep(time.Until(twentyHoursLater))
		}

		// This code block will only run once every 20 hours
		fmt.Println("Running after 20 hours...")
		// Do your logic here
	}
}
