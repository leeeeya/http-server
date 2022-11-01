package handlers

import (
	"encoding/json"
	"fmt"
	"http-server/internal/crud"
	"http-server/internal/storage"
	"log"
	"net/http"
)

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {

		var event storage.Event

		event.Date = "none"
		event.Event = "none"

		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad Request")
			log.Println(err)
			return
		}
		defer r.Body.Close()

		if err := event.UpdateValidation(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad Request")
			log.Println(err)
			return
		}

		if err := crud.UpdateData(event); err != nil {
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
