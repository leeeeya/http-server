package handlers

import (
	"fmt"
	"http-server/crud"
	"net/http"
)

func GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if data, err := crud.FetchData(month); err != nil {
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
