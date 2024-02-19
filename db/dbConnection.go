package db

import (
	"database/sql"

	_ "github.com/lib/pq"

	"log"
)

//code to connect to database
type Database struct {
	Db *sql.DB
}
func NewDb () *Database {
	return &Database{}
}

func (d *Database) Connect() error {
	var err error
	database, err := sql.Open("postgres", "postgres://postgres:Dsquare142@localhost/realestate?sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return err
	}
	d.Db = database
	return nil
}