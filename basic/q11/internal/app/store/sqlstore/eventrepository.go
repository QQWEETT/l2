package sqlstore

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"q11/internal/app/model"
	"q11/internal/app/store"
	"time"
)

type EventRepository struct {
	store *Store
}

func (r *EventRepository) Create(e *model.Event) error {
	if err := e.Validate(); err != nil {
		return err
	}
	return r.store.db.QueryRow("INSERT INTO events (user_id, name, date) VALUES ($1, $2, $3) RETURNING event_id",
		e.UID,
		e.Name,
		e.Date,
	).Scan(&e.EID)

}

func (r *EventRepository) UpdateEvent(e *model.Event) error {
	u := &model.Event{}
	return r.store.db.QueryRow(
		"UPDATE events set name=$1 WHERE (user_id = $2 AND date = $3) RETURNING (user_id)  ",
		e.Name,
		e.UID,
		e.Date,
	).Scan(
		&u.UID,
	)

}

func (r *EventRepository) DeleteEvent(e *model.Event) error {
	_, err := r.store.db.Exec(
		"DELETE FROM events WHERE (user_id = $1 AND date = $2)",
		e.UID,
		e.Date,
	)

	return err
}

func (r *EventRepository) FindByDate(user_id int, date string) (*model.Event, error) {
	u := &model.Event{}

	if err := r.store.db.QueryRow(
		"SELECT event_id, user_id, name, date FROM events WHERE (user_id = $1 AND date = $2) ",
		user_id,
		date,
	).Scan(
		&u.EID,
		&u.UID,
		&u.Name,
		&u.Date,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	fmt.Println(u)
	return u, nil

}

func (r *EventRepository) FindDateByInterval(user_id int, date string, maxdate int) (model.Events, error) {
	dateStamp, _ := time.Parse("2006-01-02", date)
	s := dateStamp.AddDate(0, 0, maxdate)
	dateString := s.Format("2006-01-02")

	rows, err := r.store.db.Query(
		"SELECT event_id, user_id, name, date FROM events WHERE (user_id = $1 AND date >= $2 and date <= $3) ",
		user_id,
		date,
		dateString)
	if err != nil {
		panic(err)
	}
	events := model.Events{}
	for rows.Next() {
		e := &model.Event{}
		if err := rows.Scan(&e.EID, &e.UID, &e.Name, &e.Date); err != nil {
			fmt.Println(err)
			return nil, model.ErrNotFoundRecord
		}

		events = append(events, e)
	}

	return events, nil

}
