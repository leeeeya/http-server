package error_handler

import (
	"log"
	"net/http"
)

// Handle обработчик ошибок сервера
func Handle(w http.ResponseWriter, code int, err error) {
	http.Error(w, http.StatusText(code), code)
	log.Println(err)
}
