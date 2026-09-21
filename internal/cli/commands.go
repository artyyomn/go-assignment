package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/artyyomn/go-assignment/internal/user"
)

// TODO:

// before login
// NOTE: after login show:Username,Registration Time, MFA stat,
// session time left, Last logged in
func (c *CLI) Login() {
	username, err := c.readUsername()
	if err != nil {
		fmt.Println("Error reading username:", err)
		return
	}

	password, err := c.readPassword()
	if err != nil {
		fmt.Println("Error reading password:", err)
		return
	}

	loggedInUser, err := c.authService.Login(context.Background(), username, password)
	if err != nil {
		fmt.Println("Login failed:", err)
		return
	}

	fmt.Println("Login successful.")
	c.showUserDetails(loggedInUser)
}

func (c *CLI) showUserDetails(loggedInUser *user.User) {
	lastLogin := "Never"
	if loggedInUser.LastLoginAt != nil {
		lastLogin = loggedInUser.LastLoginAt.Format(time.RFC3339)
	}

	fmt.Println("Username:", loggedInUser.Username)
	fmt.Println("Registration Time:", loggedInUser.CreatedAt.Format(time.RFC3339))
	fmt.Println("MFA Status: Disabled")
	fmt.Println("Session Time Left: Not available")
	fmt.Println("Last Logged In:", lastLogin)
}

func (c *CLI) Register() {
	username, err := c.readUsername()
	if err != nil {
		fmt.Println("Error reading username:", err)
		return
	}

	password, err := c.readPassword()
	if err != nil {
		fmt.Println("Error reading password:", err)
		return
	}

	confirmPassword, err := c.readPasswordConfirmation()
	if err != nil {
		fmt.Println("Error reading password:", err)
		return
	}

	if password != confirmPassword {
		fmt.Println("Passwords do not match.")
		return
	}

	err = c.authService.Register(
		context.Background(),
		username,
		password,
	)

	if err != nil {
		fmt.Println("Registration failed:", err)
		return
	}
	fmt.Println("Registration successful.")
}

func Exit() {}
func Help() {}

// after login
func Whoami()   {}
func Logout()   {}
func HelpUser() {}

// TOTP auth????
func Enable2FA()  {}
func Disable2FA() {}
