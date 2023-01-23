package databasehandler

import (
	"log"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetDataBase() *sql.DB {
	db, err := sql.Open("sqlite3", "passwords.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateTable() {
	statement, err := GetDataBase().Prepare(`CREATE TABLE IF NOT EXISTS passwords ("pid" INTEGER PRIMARY KEY AUTOINCREMENT, "password" TEXT, "website" TEXT);`)
	if err != nil {
		log.Fatal(err)
	}

	statement.Exec()
}
