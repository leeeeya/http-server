package crud

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"http-server/db"
	"http-server/internal/storage"
	"log"
)

func serialising(event []storage.Event) ([]byte, error) {
	return json.Marshal(event)
}

func FetchData(interval string) ([]byte, error) {
	var tmp storage.Event

	event := make([]storage.Event, 0)
	query := fmt.Sprintf("SELECT user_id, date, event FROM events WHERE date >= now() - '1%s'::interval", interval)

	if rows, err := db.DB.Query(query); err == nil {
		defer rows.Close()
		for rows.Next() == true {
			if err := rows.Scan(&tmp.UserID, &tmp.Date, &tmp.Event); err != nil {
				log.Println(err)
				return nil, err
			} else {
				event = append(event, tmp)
			}
		}
	} else if err != sql.ErrNoRows {
		log.Println(err)
		return nil, err
	}
	return serialising(event)
}
