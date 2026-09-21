package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/artyyomn/go-assignment/internal/auth"
	"github.com/chzyer/readline"
)

type CLI struct {
	authService *auth.Service
	session     *auth.Session
	rl          *readline.Instance
}

func NewCLI(authService *auth.Service) *CLI {
	return &CLI{
		authService: authService,
	}
}

func (c *CLI) RunCLI() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "~> ",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	c.rl = rl
	if err != nil {
		log.Fatal("Error geting readline module")
	}
	defer rl.Close()

	// Welcome message looks ugly, fix later
	fmt.Println("Welcome to Interactive Go-User CLI")
	fmt.Println("usage:	 /[command] [options]")
	fmt.Println("commands:")
	fmt.Println("	 /help")
	fmt.Println("	 /register")
	fmt.Println("	 /login")
	fmt.Println("	 /exit")

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
			c.Login()
		case inputs[0] == "/register":
			c.Register()
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

func (c *CLI) readUsername() (string, error) {
	c.rl.SetPrompt("Username: ")

	username, err := c.rl.Readline()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(username), nil
}

func (c *CLI) readPassword() (string, error) {
	c.rl.SetPrompt("Password: ")

	setPasswordCfg := c.rl.GenPasswordConfig()
	setPasswordCfg.SetListener(func(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
		c.rl.SetPrompt(fmt.Sprintf("Enter password(%v): ", len(line)))
		c.rl.Refresh()
		return nil, 0, false
	})

	password, err := c.rl.ReadPasswordWithConfig(setPasswordCfg)
	if err != nil {
		return "", err
	}

	return string(password), nil
}

func (c *CLI) readPasswordConfirmation() (string, error) {
	c.rl.SetPrompt("Confirm password: ")

	setPasswordCfg := c.rl.GenPasswordConfig()
	setPasswordCfg.SetListener(func(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
		c.rl.SetPrompt(fmt.Sprintf("Enter password(%v): ", len(line)))
		c.rl.Refresh()
		return nil, 0, false
	})

	password, err := c.rl.ReadPasswordWithConfig(setPasswordCfg)
	if err != nil {
		return "", err
	}

	return string(password), nil
}
