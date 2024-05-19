package database

import (
	"KaifTravel/utils"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitSqliteDB() *sql.DB {
	var err error

	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateDB() {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`)
	if err != nil {
		panic(err.Error())
	}

	isEmpty, err := IsTableEmpty("users")
	if err != nil {
		panic(err.Error())
	}
	if isEmpty {
		CreateUser("admin", utils.RandomPassword(64))
	}
}

func IsTableEmpty(tableName string) (bool, error) {
	var rowCount int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s;", tableName)
	err := db.QueryRow(query).Scan(&rowCount)
	if err != nil {
		return false, err
	}
	return rowCount == 0, nil
}
