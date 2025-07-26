package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/InfiniteCodeMonkeys/weather-cli/services"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Must provide a command and a city name.")
		return
	}

	command := args[1]
	city := strings.Join(os.Args[2:], " ")

	fmt.Println("Command:", command)
	fmt.Println("City:", city)

	location, err := services.GetLocation(city)
	if err != nil {
		fmt.Println("Error fetching location:", err)
		return
	}

	switch command {
	case "get":
		fmt.Println("Fetching today's weather for:", city)
		// Call the GetWeather function from services package

		services.GetWeather(location.Lat, location.Lon, false)
		// Your logic here
	case "forecast":
		fmt.Println("Fetching weather forecast for:", city)
		// Your logic here
		services.GetWeather(location.Lat, location.Lon, true)
	default:
		fmt.Println("Unknown command:", command)
	}
}
