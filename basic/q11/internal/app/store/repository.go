package store

import (
	"q11/internal/app/model"
)

type EventRepository interface {
	Create(event *model.Event) error
	FindByDate(int, string) (*model.Event, error)
	UpdateEvent(event *model.Event) error
	DeleteEvent(event *model.Event) error
	FindDateByInterval(int, string, int) (model.Events, error)
}
