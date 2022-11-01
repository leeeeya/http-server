package handlers

import (
	"fmt"
	"http-server/internal/crud"
	"http-server/internal/storage"
	"http-server/internal/validation"
	"log"
	"net/http"
)

func GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		uid, date, err := validation.ParseQueryStr(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad Request")
			log.Println(err)
			return
		}
		if data, err := crud.FetchData(storage.Week, date, uid); err != nil {
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
