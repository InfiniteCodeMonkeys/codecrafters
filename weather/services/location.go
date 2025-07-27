package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/InfiniteCodeMonkeys/weather/utils"
)

type Location struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func GetLocation(city string) (*Location, error) {
	fmt.Println("Fetching location for", city)

	encodedCity := url.QueryEscape(city)

	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", encodedCity)
	resp, err := utils.Fetch(url, nil)
	if err != nil {
		fmt.Println("Error fetching location:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Process the response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var locations []Location
	if err := json.Unmarshal(responseBody, &locations); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return nil, err
	}
	if len(locations) == 0 {
		fmt.Println("No location found for", city)
		return nil, fmt.Errorf("no location found")
	}
	return &locations[0], nil
}
