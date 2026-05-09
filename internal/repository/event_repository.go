package repository

import (
	"database/sql"
	"errors"

	"exammple.com/event-booking-api/internal/domain"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) Save(event *domain.Event) error {
	query := `INSERT INTO events(
		name, description, location, dateTime, user_id)
		VALUES (?, ?, ?, ?, ?)`

	stmt, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	event.ID = id

	return nil
}

func (r *EventRepository) GetAll() ([]domain.Event, error) {
	query := `SELECT * FROM events`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	var events []domain.Event

	for rows.Next() {
		var event domain.Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	defer rows.Close()

	return events, nil
}

func (r *EventRepository) GetById(id string) (*domain.Event, error) {
	var event domain.Event

	row := r.db.QueryRow("SELECT * FROM events WHERE id = $1", id)

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, errors.New("No event found with this id")
	}

	return &event, nil
}

func (r *EventRepository) Delete(id string) error {
	query := `DELETE FROM events WHERE id = ?`

	_, err := r.db.Exec(query, id)

	return err
}

func (r *EventRepository) Update(id string, data domain.Event) error {
	query := `UPDATE events SET
	name = ?, description = ?, location = ?, dateTime = ?
	WHERE ID = ?`

	stmt, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(data.Name, data.Description, data.Location, data.DateTime, id)

	return err
}

func (r *EventRepository) RegisterUserForEvent(event domain.Event, userId int64) error {
	query := `INSERT INTO registrations(
		user_id, event_id)
		VALUES (?, ?)`

	stmt, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId, event.ID)

	return err
}

func (r *EventRepository) GetRegistrationsForEvent(event domain.Event) ([]int64, error) {
	query := `SELECT user_id FROM registrations WHERE event_id = ?`

	rows, err := r.db.Query(query, event.ID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var userIds []int64

	for rows.Next() {
		var userId int64
		if err := rows.Scan(&userId); err != nil {
			return nil, err
		}
		userIds = append(userIds, userId)
	}

	return userIds, nil
}

func (r *EventRepository) CancelUserRegistrationForEvent(event domain.Event, userId int64) error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`

	stmt, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID, userId)

	return err
}
