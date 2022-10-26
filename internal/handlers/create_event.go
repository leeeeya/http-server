package handlers

import (
	"encoding/json"
	"fmt"
	"http-server/crud"
	"http-server/internal/storage"
	"log"
	"net/http"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		var event storage.Event

		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad Request")
			log.Println(err)
			return
		}
		defer r.Body.Close()

		if err := event.CreateValidation(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad Request")
			log.Println(err)
			return
		}

		if err := crud.InsertData(event); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "Service Unavailable")
			log.Println("create_event: Service Unavailable")
			return
		}
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Service Unavailable")
		log.Println("create_event: Service Unavailable")
	}
}
