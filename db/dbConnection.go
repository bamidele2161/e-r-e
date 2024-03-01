package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"log"
)

//code to connect to database
type Database struct {
	Db *sql.DB
}
func NewDb () (*Database, error){
	db := &Database{}
	err := db.Connect()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *Database) Connect() error {
	var err error
	database, err := sql.Open("postgres", "postgres://postgres:Dsquare142@localhost/realestate?sslmode=disable")
	if err != nil {
		fmt.Println("checking error",err)
		log.Fatal(err)
		return err
	}
	d.Db = database
	return nil
}