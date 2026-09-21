package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/artyyomn/go-assignment/internal/auth"
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

	loggedInUser, session, err := c.authService.Login(context.Background(), username, password)
	if err != nil {
		fmt.Println("========================================")
		//colorred text rendering codes
		fmt.Println("\033[1;31mLogin failed:\033[0m", err)
		return
	}

	fmt.Println("========================================")
	fmt.Println("\033[1;32mLogin successful.\033[0m")
	c.currentUser = loggedInUser
	c.session = session
	c.showUserDetails(loggedInUser, session)
}

func (c *CLI) showUserDetails(loggedInUser *user.User, session *auth.Session) {
	lastLogin := "Never"
	if loggedInUser.LastLoginAt != nil {
		lastLogin = loggedInUser.LastLoginAt.Format(time.RFC3339)
	}

	fmt.Println("Username:", loggedInUser.Username)
	fmt.Println("Registration Time:", loggedInUser.CreatedAt.Format(time.RFC3339))
	fmt.Println("MFA Status: Disabled")
	//fmt.Println("Session ID:", session.ID)
	fmt.Println("Session Time Left:", time.Until(session.ExpiresAt).Round(time.Second))
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
		fmt.Println("\033[1;31mRegistration failed:\033[0m", err)
		return
	}
	fmt.Println("\033[1;32mRegistration successful.\033[0m")
}

func Exit() {}

func (c *CLI) Help() {
	fmt.Println("usage: /[command]")
	fmt.Println("commands:")
	fmt.Println("  /help")
	fmt.Println("  /register")
	fmt.Println("  /login")
	fmt.Println("  /exit")
	fmt.Println("  /whoami")
	fmt.Println("  /logout")
}

// after login
func (c *CLI) Whoami() {
	if !c.requireLogin() {
		return
	}
	fmt.Println("Username:", c.currentUser.Username)
}

func (c *CLI) Logout() {
	if !c.requireLogin() {
		return
	}

	if err := c.authService.Logout(context.Background(), c.session.ID); err != nil {
		fmt.Println("\033[1;31mLogout failed:\033[0m", err)
		return
	}

	c.currentUser = nil
	c.session = nil
	fmt.Println("\033[1;32mLogout successful.\033[0m")
	//fmt.Println("Logout successful.")
}

func (c *CLI) isLoggedIn() bool {
	if c.currentUser == nil || c.session == nil {
		return false
	}
	if !time.Now().Before(c.session.ExpiresAt) {
		c.currentUser = nil
		c.session = nil
		return false
	}
	return true
}

func (c *CLI) requireLogin() bool {
	if c.isLoggedIn() {
		return true
	}
	fmt.Println("You must be logged in to use this command.")
	return false
}

// TOTP auth????
func Enable2FA()  {}
func Disable2FA() {}
