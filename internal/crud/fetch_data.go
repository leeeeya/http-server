package crud

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"http-server/db"
	"http-server/internal/storage"
	"log"
)

// сериализация данных для отправки клиенту
func serialising(event []storage.Event) ([]byte, error) {
	return json.Marshal(event)
}

// получение параметров из queryString для метода GET
func handleQueryString(interval, date string, uid int) (string, string) {
	var uidCond, dateCond string

	if uid == 0 {
		uidCond = ""
	} else {
		uidCond = fmt.Sprintf("user_id = %d ", uid)
	}

	if date == "" {
		switch interval {
		case storage.Day:
			dateCond = "date = date(now())"
		default:
			dateCond = fmt.Sprintf("date <= now() - '1%s'::interval", interval)
		}
	} else {
		switch interval {
		case storage.Day:
			dateCond = fmt.Sprintf("date = '%s'", date)
		default:
			dateCond = fmt.Sprintf("date between '%s' and date('%s') + '1%s'::interval", date, date, interval)
		}
	}
	return uidCond, dateCond
}

// FetchData получение данных из БД
func FetchData(interval, date string, uid int) ([]byte, error) {
	var tmp storage.Event
	var and string

	event := make([]storage.Event, 0)

	uidCond, dateCond := handleQueryString(interval, date, uid)

	if uidCond != "" {
		and = "and "
	}
	query := "SELECT * FROM events WHERE " + uidCond + and + dateCond

	if rows, err := db.DB.Query(query); err == nil {
		defer rows.Close()
		for rows.Next() == true {
			if err := rows.Scan(&tmp.ID, &tmp.UserID, &tmp.Date, &tmp.Event); err != nil {
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
