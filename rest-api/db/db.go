package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var logger = log.Default()
var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
	logger.Println("Successfully connected to database")
}

func createTables() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    name TEXT NOT NULL,
		    description TEXT NOT NULL,
		    location TEXT NOT NULL,
		    dateTime DATE NOT NULL,
		    user_id INTEGER
		)
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("could not create events table")
	}
}
