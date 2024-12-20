package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var logger = log.Default()

var DB *sql.DB

func InitDB() {
	setupDB("sqlite3", "api.db")

	createTables()
	logger.Println("Successfully connected to database")
}

func setupDB(driverName, sourceName string) {
	var err error
	DB, err = sql.Open(driverName, sourceName)

	if err != nil {
		panic("could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    email TEXT NOT NULL UNIQUE,
		    password TEXT NOT NULL
		)
	`
	createTable("users", createUsersTable)

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    name TEXT NOT NULL,
		    description TEXT NOT NULL,
		    location TEXT NOT NULL,
		    dateTime DATE NOT NULL,
		    user_id INTEGER,
		    FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`
	createTable("events", createEventsTable)
}

func createTable(tableName, query string) {
	_, err := DB.Exec(query)

	if err != nil {
		panic("could not create " + tableName + " table")
	}
}
