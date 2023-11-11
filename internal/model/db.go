package model

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Setup() {
	name := "db/mydb.db"
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

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()
	for i := 0; i < 25; i++ {
		tx.Exec("INSERT INTO Todo (description, is_done) VALUES (?,?)", "hello", i > 0)
	}

}

func Close() {
	err := db.Close()
	log.Fatal(err)
}
