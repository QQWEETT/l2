package model

import (
	"errors"
)

var (
	errInvalidUserID = errors.New("user_id should be positive integer")
	errInvalidDate   = errors.New("date should be in YYYY-MM-DD format")
	errInvalidInfo   = errors.New("info field is required and should has min 3 symbols")
)

type Event struct {
	EID  int    `json:"event_id"`
	UID  int    `json:"user_id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

type Events []*Event

func (e *Event) Validate() error {
	if e.UID <= 0 {
		return errInvalidUserID
	}

	return nil
}

var ErrNotFoundRecord = errors.New("события не найдены")
