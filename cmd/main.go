package main

import (
	"log"

	"github.com/artyyomn/go-assignment/internal/auth"
	"github.com/artyyomn/go-assignment/internal/cli"
	"github.com/artyyomn/go-assignment/internal/db"
	user "github.com/artyyomn/go-assignment/internal/user/adapters"
)

func main() {

	db, err := db.NewDB("data/app.db")
	if err != nil{
		log.Fatal("error conneting go db", err)
	}

	userRepo := user.NewSQLiteRepository(db)
	authService := auth.NewService(userRepo)

	shell := cli.NewCLI(authService)

	shell.RunCLI()
}
