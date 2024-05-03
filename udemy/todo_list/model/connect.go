package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mysql")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the databse")
	con = db
	return db, nil
}
