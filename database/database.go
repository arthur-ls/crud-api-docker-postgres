package database

import (
	"database/sql"
	"fmt"
)

func InitDB() (*sql.DB, error) {
	host, port, user, password, dbname := DatabaseParams()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db, err
}
