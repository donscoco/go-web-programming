package data

import (
	"database/sql"
	"log"
)

//  db pointer
var Db *sql.DB

//init
func init() {
	// os connect to sql
	Db, err := sql.Open("postgres", "dbname=chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)

	}
	return
}
