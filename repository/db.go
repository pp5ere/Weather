package repository

import (
	"database/sql"
)

//SqliteDB defines the connection with database
type SqliteDB struct{
	drive string
	tp string
	connection *sql.DB
}

//New creates a new Repository layer sharing the sqlite connection
func New(drive string, tp string) (*SqliteDB, error) {
	//db, err := sql.Open("sqlite3", "./db.sqlite3")
	db, err := sql.Open(drive, tp)
	return &SqliteDB{
		drive 		: drive,
		tp			: tp,
		connection	: db,
	}, err
}
