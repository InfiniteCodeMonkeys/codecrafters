package handlers

import (
	"fmt"
	"net/http"
)

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}
