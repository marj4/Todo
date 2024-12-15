package db

import (
	"database/sql"
	"fmt"
)

const (
	ErrorConnect = "Cant connect to db. Error: %v"
)

func Connect(DatabaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", DatabaseUrl)
	if err != nil {
		return nil, fmt.Errorf(ErrorConnect, err)
	}

	defer db.Close()

	return db, nil
}
