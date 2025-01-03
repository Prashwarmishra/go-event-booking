package models

import (
	"errors"
	"go-event-booking/db"
	"go-event-booking/utils"
)

type User struct {
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	
	hashedPassword, err := utils.HashPassword(u.Password)

	if err!= nil {
        return err
    }
	
	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil{
		return err
	}

	u.ID = id
	return err
}

func (u *User) ValidateUser() error {
	query := `SELECT password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)
	var hashedPassword string
	err := row.Scan(&hashedPassword)

	if err != nil {
		return err
	}

	if !utils.ValidatePassword(hashedPassword, u.Password) {
		return errors.New("Invalid password")
	}

	return err
}
