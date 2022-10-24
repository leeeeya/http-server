package main

import (
	"http-server/db"
	"http-server/internal/handlers"
	"log"
	"net/http"
)

func main() {
	cfg := InitConfig()

	db.InitDB(cfg.DatabaseConn)
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/calendar/create_event", handlers.CreateEvent)

	mux.HandleFunc("/calendar/events_for_day", handlers.GetEventsForDay)
	mux.HandleFunc("/calendar/events_for_week", handlers.GetEventsForWeek)
	mux.HandleFunc("/calendar/events_for_month", handlers.GetEventsForMonth)

	err := http.ListenAndServe(":"+cfg.Port, mux)
	if err != nil {
		log.Fatalln(err)
	}
}
