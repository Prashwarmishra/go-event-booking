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
	UserID int
}

var events []Event

func (e Event) Save() (Event, error) {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Error preparing query", err)
		return Event{}, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		fmt.Println("Error executing query", err)
		return Event{}, err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return e, err
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