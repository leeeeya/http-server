package validation

import (
	"net/http"
	"strconv"
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
	return uid, date, nil
}
