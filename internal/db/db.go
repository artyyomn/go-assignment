package db

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

func NewDB(path string)(*sql.DB, error){
	db, err := sql.Open("sqlite", path)
	if err != nil{
		log.Println("Error opening database", err)
		return nil, err
	}
	
	//TODO:
	// Add automatic migrations

	//Ping the database for sanity
	err = db.Ping()
	if err != nil{
		log.Fatal("Error pinging database", err)
	}

	return db, nil
}
