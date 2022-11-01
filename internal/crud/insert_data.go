package crud

import (
	"fmt"
	"http-server/db"
	"http-server/internal/storage"
	_ "http-server/internal/storage"
)

// InsertData внесение данных, полученных от клиента, в БД
func InsertData(event storage.Event) error {

	query := fmt.Sprintf("INSERT INTO events(user_id, date, event) VALUES (%d, '%s', '%s') on conflict do nothing",
		event.UserID, event.Date, event.Event)

	if _, err := db.DB.Exec(query); err != nil {
		return err
	}
	return nil
}
