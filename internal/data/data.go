package data

import "database/sql"

type DataStore struct {
	db *sql.DB
}

func NewDataStore(db *sql.DB) *DataStore {
	return &DataStore{db: db}
}
