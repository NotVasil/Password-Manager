package db

import (
	"log"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetDataBase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "passwords.db")
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

func CreateTable() error {
	db, err := GetDataBase()
	statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS passwords ("pid" INTEGER PRIMARY KEY AUTOINCREMENT, "password" TEXT, "website" TEXT);`)

	if err != nil {
		log.Fatal(err)
	}

	statement.Exec()
	return err
}
