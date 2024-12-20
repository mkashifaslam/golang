package models

import (
	"errors"
	"github.com/mkashifaslam/golang/rest-api/db"
	"github.com/mkashifaslam/golang/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id

	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	logger.Printf("\nUserEmail: %s\n", u.Email)
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	var retrievedID int64
	err := row.Scan(&retrievedID, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	logger.Printf("\nUserID: %d\n", retrievedID)
	logger.Printf("HashPassword: %s\nPassword: %s\n", retrievedPassword, u.Password)

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	logger.Println("PasswordValid", passwordIsValid)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	u.ID = retrievedID

	return nil
}
