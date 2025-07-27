package handlers

import (
	"net/http"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

}
