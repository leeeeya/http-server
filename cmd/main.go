package main

import (
	"http-server/db"
	"http-server/internal/config"
	"http-server/internal/handlers"
	"http-server/internal/middleware"
	"log"
	"net/http"
)

func main() {
	cfg := config.InitConfig()

	db.InitDB(cfg.DatabaseConn)
	defer db.Close()

	mux := http.NewServeMux()
	handler := middleware.Logging(mux)

	mux.HandleFunc("/calendar/create_event", handlers.CreateEvent)
	mux.HandleFunc("/calendar/update_event", handlers.UpdateEvent)
	mux.HandleFunc("/calendar/delete_event", handlers.DeleteEvent)

	mux.HandleFunc("/calendar/events_for_day", handlers.GetEventsForDay)
	mux.HandleFunc("/calendar/events_for_week", handlers.GetEventsForWeek)
	mux.HandleFunc("/calendar/events_for_month", handlers.GetEventsForMonth)

	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))
}
