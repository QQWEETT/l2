package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"q11/internal/app/store"
)

type Store struct {
	db              *sql.DB
	eventRepository *EventRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Event() store.EventRepository {
	if s.eventRepository != nil {
		return s.eventRepository
	}
	s.eventRepository = &EventRepository{
		store: s,
	}
	return s.eventRepository
}
