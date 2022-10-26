package crud

import (
	"fmt"
	"http-server/db"
	"http-server/internal/storage"
	"log"
)

func UpdateData(e storage.Event) error {
	var updDate, updEvent string
	comma := ","

	if e.Date != "" {
		updDate = fmt.Sprintf("date = '%s'", e.Date)
	}

	if e.Event != "" {
		updEvent = fmt.Sprintf("event = '%s'", e.Event)
	} else {
		comma = ""
	}

	query := fmt.Sprintf("UPDATE events SET %s %s %s WHERE id = %d", updDate, comma, updEvent, e.ID)
	if _, err := db.DB.Exec(query); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
