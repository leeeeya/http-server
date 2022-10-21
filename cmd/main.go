package main

import (
	"log"
	"net/http"
	"time"

	"http-server/internal/handlers"
)

type Event struct {
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
	Event  string    `json:"event"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/events_for_day", handlers.GetEventsForDay)
	
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln(err)
	}
}
