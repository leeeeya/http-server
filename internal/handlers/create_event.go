package handlers

import (
	"encoding/json"
	"http-server/internal/crud"
	error_handler "http-server/internal/error-handler"
	"http-server/internal/storage"
	"log"
	"net/http"
)

// CreateEvent обработчик для создания записи в календаре
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		var event storage.Event

		// десериализация данных, полученных в POST-запросе, для записи в БД
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			error_handler.Handle(w, http.StatusInternalServerError, err)
			return
		}
		defer r.Body.Close()

		if err := event.CreateValidation(); err != nil {
			error_handler.Handle(w, http.StatusBadRequest, err)
			return
		}

		if err := crud.InsertData(event); err != nil {
			error_handler.Handle(w, http.StatusBadRequest, err)
			return
		}
	} else {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(http.StatusText(http.StatusInternalServerError))
	}
}
