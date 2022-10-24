package storage

type Event struct {
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
	Event  string `json:"event"`
}
