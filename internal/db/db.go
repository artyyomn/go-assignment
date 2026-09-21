package db

import (
	"database/sql"
	"log"
)

func NewDB(path string)(*sql.DB, error){
	db, err := sql.Open("sqlite", path)
	if err != nil{
		log.Println("Error opening database", err)
		return nil, err
	}

	//Ping the database for sanity
	err = db.Ping()
	if err != nil{
		log.Fatal("Error pinging database", err)
	}

	return db, nil
}
