package storage

import (
	"fmt"
	"time"
)

type Event struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
	Event  string `json:"event"`
}

func (e Event) Validation() error {
	if err := e.DateValidation(); err != nil {
		return err
	}
	if err := e.UidValidation(); err != nil {
		return err
	}
	if err := e.EventValidation(); err != nil {
		return err
	}
	return nil
}

func (e *Event) DateValidation() error {
	_, err := time.Parse("2006-01-02", e.Date)
	return err
}

func (e Event) UidValidation() error {
	if e.UserID < 1 {
		return fmt.Errorf("invalid user_id")
	}
	return nil
}

func (e Event) EventValidation() error {
	if e.Event == "" {
		return fmt.Errorf("invalid event")
	}
	return nil
}

func (e Event) IdValidation() error {
	if e.ID < 1 {
		return fmt.Errorf("invalid ID")
	}
	return nil
}
