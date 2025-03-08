package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Input struct {
	House string `json:"house"`
	Age float32 `json:"age"`
}

type Output struct {
	House string `json:"house"`
	AverageAge int `json:"average_age"`
}

type HouseStats struct {
    Sum   int
    Count int
}

func server() (string, error) {
	resp, err := http.Get("https://coderbyte.com/api/challenges/json/wizard-list")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result []Input
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	
	houseData := make(map[string]*HouseStats)

	for _, h := range result {
		if h.House == "" {
			continue
		}

		if _, exists := houseData[h.House]; !exists {
			houseData[h.House] = &HouseStats{}
		}

		houseData[h.House].Sum += int(h.Age)
		houseData[h.House].Count++
	}

	var output []Output
	for house, data := range houseData {
		averageAge := data.Sum / data.Count
		output = append(output, Output{
			House: house,
			AverageAge: averageAge,
		})
	}

	// Pretty-print the JSON response
	prettyJSON, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return "", err
	}

	return string(prettyJSON), nil
}

func main() {
	output, err := server()
	if err != nil {
		log.Fatal("Error formatting JSON:", err)
	}
	fmt.Println(output)
}