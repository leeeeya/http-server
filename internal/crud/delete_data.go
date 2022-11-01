package crud

import (
	"http-server/db"
	"log"
)

// DeleteData удаление записей из БД
func DeleteData(id int) error {
	if _, err := db.DB.Exec(`DELETE FROM events.public.events WHERE id = $1`, id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
