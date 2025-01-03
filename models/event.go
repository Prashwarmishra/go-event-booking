package models

import (
	"fmt"
	"go-event-booking/db"
	"time"
)

type Event struct {
	ID int64
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int64
}

var events []Event

func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Error preparing query", err)
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		fmt.Println("Error executing query", err)
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	var event Event
	
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}
	
	return &event, err
}

func (e Event) UpdateEvent() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, datetime = ?
	WHERE id = ?
	`
	
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	return err
}

func (e Event) DeleteEvent () error {
	script := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(script)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}

func (e Event) GetAllRegistrations() ([]Registation, error) {
	script := `SELECT * FROM registrations WHERE event_id = ?`

	rows, err := db.DB.Query(script, e.ID)
	registrations := []Registation{}
	
	if err != nil {
		return []Registation{}, err
	}

	defer rows.Close()
	
	for rows.Next() {
		var registration Registation
		rows.Scan(&registration.ID, &registration.EventId, &registration.UserId)
		registrations = append(registrations, registration)
	}
	
	return registrations, err
} 

func (e *Event) CreateRegistration(userId int64) error {
	script := `INSERT INTO registrations (event_id, user_id) values (?, ?)`
	stmt, err := db.DB.Prepare(script)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e *Event) CancelRegistration(userId int64) error {
	script := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`

	_, err := db.DB.Exec(script, e.ID, userId)

	if err != nil {
		return err
	}

	return err
}