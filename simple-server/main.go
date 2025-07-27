package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/InfiniteCodeMonkeys/simple-server/handlers"
	"github.com/InfiniteCodeMonkeys/simple-server/middleware"
)

func main() {
	// Serve static assets (CSS, JS, images) from /static/
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", middleware.Logging(handlers.WelcomeHandler))
	http.HandleFunc("/greet", middleware.Logging(handlers.GreetHandler))
	http.HandleFunc("/weather", middleware.Logging(handlers.WeatherHandler))
	http.HandleFunc("/api/message", middleware.Logging(handlers.MessageHandler))
	http.HandleFunc("/api/echo", middleware.Logging(handlers.EchoHandler))

	fmt.Println("Server is running on port 8080")

	srv := &http.Server{Addr: ":8080"}
	go srv.ListenAndServe()

	// listen for Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	srv.Shutdown(context.Background())

	go func() {
		log.Println("Background task running...")
	}()
}
