package main

import (
	"log"
	"time"

	"github.com/artyyomn/go-assignment/internal/auth"
	authadapter "github.com/artyyomn/go-assignment/internal/auth/adapters"
	"github.com/artyyomn/go-assignment/internal/cli"
	"github.com/artyyomn/go-assignment/internal/db"
	user "github.com/artyyomn/go-assignment/internal/user/adapters"
)

func main() {

	db, err := db.NewDB("data/app.db")
	if err != nil {
		log.Fatal("error conneting go db", err)
	}

	userRepo := user.NewSQLiteRepository(db)
	sessionRepo := authadapter.NewSQLiteSessionRepository(db)
	authService := auth.NewService(userRepo, sessionRepo, auth.Config{
		SessionTimeout: 2 * time.Hour,
		// 2 hours ?????
	})

	shell := cli.NewCLI(authService)

	shell.RunCLI()
}
