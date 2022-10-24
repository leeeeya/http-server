package handlers

import (
	"fmt"
	"http-server/crud"
	"net/http"
)

func GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if data, err := crud.FetchData(week); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal Server Error")
		} else {
			w.Write(data)
		}
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Service Unavailable")
	}
}
