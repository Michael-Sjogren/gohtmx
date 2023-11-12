package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Setup() {
	name := "mydb.db"
	os.Remove(name)
	var err error
	db, err = sql.Open("sqlite3", name)
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(`CREATE TABLE IF NOT EXISTS Todo ( 
		id INTEGER PRIMARY KEY,
		description TEXT NOT NULL,
		is_done INTEGER NOT NULL DEFAULT 0 )`, nil)

	if err != nil {
		log.Fatal(err)
	}

}

func Close() {
	err := db.Close()
	log.Fatal(err)
}
