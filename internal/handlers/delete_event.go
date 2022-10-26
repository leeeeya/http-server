package handlers

import (
	"encoding/json"
	"fmt"
	"http-server/crud"
	"log"
	"net/http"
)

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		type Id struct {
			Id int `json:"id"`
		}
		var id Id

		if err := json.NewDecoder(r.Body).Decode(&id.Id); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad Request")
			log.Println(err)
			return
		}
		defer r.Body.Close()

		if id.Id < 1 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad Request")
			log.Println("invalid id")
			return
		}

		if err := crud.DeleteData(id.Id); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "Service Unavailable")
			log.Println("create_event: Service Unavailable")
			return
		}
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Service Unavailable")
		log.Println("create_event (35): Service Unavailable")
	}
}
