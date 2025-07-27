package services

import (
	"encoding/json"
	"fmt"

	"github.com/InfiniteCodeMonkeys/weather/utils"
)

type WeatherResponse struct {
	Properties struct {
		ForecastHourly string `json:"forecastHourly"`
		Forecast       string `json:"forecast"`
	} `json:"properties"`
}

type ForecastResponse struct {
	Properties struct {
		Periods []struct {
			DetailedForecast string `json:"detailedForecast"`
			Temperature      int    `json:"temperature"`
			TemperatureUnit  string `json:"temperatureUnit"`
			Name             string `json:"name"`
		} `json:"periods"`
	} `json:"properties"`
}

func GetWeather(lat string, lon string, sevenDay bool) {

	headers := map[string]string{
		"User-Agent": "my-weather-app (mike@chindogulabs.com)",
		"Method":     "GET",
	}

	resp, err := utils.Fetch(fmt.Sprintf("https://api.weather.gov/points/%s,%s", lat, lon), headers)
	if err != nil {
		fmt.Println("Error fetching weather data:", err)
		return
	}

	defer resp.Body.Close()

	var weatherResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	forecastUrl := weatherResp.Properties.Forecast

	forecastResp, err := utils.Fetch(forecastUrl, map[string]string{"Method": "GET"})
	if err != nil {
		fmt.Println("Error fetching hourly forecast:", err)
		return
	}

	defer forecastResp.Body.Close()

	var forecast ForecastResponse
	if err := json.NewDecoder(forecastResp.Body).Decode(&forecast); err != nil {
		fmt.Println("Error decoding forecast response:", err)
		return
	}

	if sevenDay {
		for _, period := range forecast.Properties.Periods {
			fmt.Printf("%s: %s, Temperature: %d%s\n", period.Name, period.DetailedForecast, period.Temperature, period.TemperatureUnit)
			fmt.Println("\n---")
		}

	} else {
		fmt.Printf("Weather forecast for %s: %s\n", forecast.Properties.Periods[0].Name, forecast.Properties.Periods[0].DetailedForecast)

	}
}
