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

func (e Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	defer stmt.Close()

	if err != nil {
		fmt.Println("Error preparing query", err)
		return err
	}

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime)

	if err != nil {
		fmt.Println("Error executing query", err)
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() []Event {
	return events
}