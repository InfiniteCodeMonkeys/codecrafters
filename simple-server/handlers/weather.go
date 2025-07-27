package handlers

import (
	"fmt"
	"net/http"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Weather API\n")
}
