package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

)

var db *sql.DB

func open() (*sql.DB, error) {
	if db != nil  {
		return db, nil
	}
	var err error
	db, err = sql.Open("sqlite3", "./mapr.db")
	if err != nil {
		return nil, err
	}
	return db, nil	
}

func Close() error {
	return db.Close()
}

func Migrate() error {
	db, err := open()
	if err != nil {
		return err
	}

	db.Exec("CREATE TABLE Version (version varchar)")
	db.Exec("CREATE TABLE CachePR (AzureId)")
	return nil
}
