package handlers

import (
	"fmt"
	"net/http"
)

func GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "events for day")
}
