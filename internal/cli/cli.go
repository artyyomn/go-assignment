package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/artyyomn/go-assignment/internal/auth"
	"github.com/chzyer/readline"
)

type CLI struct{
	authService *auth.Service
	session *auth.Session
}

func NewCLI(authService *auth.Service) *CLI{
	return &CLI{
		authService: authService,
	}
}

func (c *CLI) RunCLI(){
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "> ",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		log.Fatal("Error geting readline module")
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		inputs := strings.Split(line, " ")

		switch {
		case inputs[0] == "/exit":
			log.Println("exited")
		case inputs[0] == "/help":
			log.Println("asked for help")
		case inputs[0] == "/login":
			log.Println("asked for login")
		case inputs[0] == "/register":
			log.Println("asked for register")
			default: 
			fmt.Println("Invalid command")
			fmt.Println("usage:	 /[command] [options]")
			fmt.Println("commands:")
			fmt.Println("	 /help")
			fmt.Println("	 /register")
			fmt.Println("	 /login")
			fmt.Println("	 /exit")
		}
	}
}
