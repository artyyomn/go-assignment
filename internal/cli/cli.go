package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/artyyomn/go-assignment/internal/auth"
	"github.com/artyyomn/go-assignment/internal/user"
	"github.com/chzyer/readline"
)

type CLI struct {
	authService *auth.Service
	currentUser *user.User
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
	fmt.Println("Type /help for available commands")
	//fmt.Println("usage:	 /[command] [options]")
	//fmt.Println("commands:")
	//fmt.Println("	 /help")
	//fmt.Println("	 /register")
	//fmt.Println("	 /login")
	//fmt.Println("	 /exit")
	//fmt.Println("	 /whoami")
	//fmt.Println("	 /logout")

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		inputs := strings.Fields(line)
		if len(inputs) == 0 {
			continue
		}

		switch {
		case inputs[0] == "/exit":
			c.Exit()
			return
		case inputs[0] == "/help":
			c.Help()
		case inputs[0] == "/clear":
			c.Clear()
		case inputs[0] == "/login":
			c.Login()
		case inputs[0] == "/register":
			c.Register()
		case inputs[0] == "/whoami":
			c.Whoami()
		case inputs[0] == "/logout":
			c.Logout()
		default:
			fmt.Println("Invalid command")
			c.Help()
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
