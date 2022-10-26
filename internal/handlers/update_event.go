package handlers

import (
	"encoding/json"
	"fmt"
	"http-server/crud"
	"http-server/internal/storage"
	"log"
	"net/http"
)

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {

		var event storage.Event

		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad Request")
			log.Println(err)
			return
		}
		defer r.Body.Close()
		fmt.Println(event)

		if err := crud.InsertData(event); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "Service Unavailable")
			log.Println("create_event (28): Service Unavailable")
			return
		}
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Service Unavailable")
		log.Println("create_event (35): Service Unavailable")
	}
}
