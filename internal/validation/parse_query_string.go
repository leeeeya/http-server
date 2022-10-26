package validation

import (
	"net/http"
	"strconv"
	"time"
)

func ParseQueryStr(r *http.Request) (int, string, error) {
	var uid int
	var err error

	uidStr := r.URL.Query().Get("user_id")
	if uidStr != "" {
		uid, err = strconv.Atoi(uidStr)
		if err != nil {
			return 0, "", err
		}
	} else {
		uid = 0
	}

	date := r.URL.Query().Get("date")
	if date != "" {
		if _, err := time.Parse("2006-01-02", date); err != nil {
			return 0, "", err
		}
	}
	return uid, date, nil
}
