package cli

import (
	"context"
	"errors"
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
	if errors.Is(err, auth.ErrTOTPRequired) {
		code, codeErr := c.readTOTPCode()
		if codeErr != nil {
			fmt.Println("Error reading authentication code:", codeErr)
			return
		}
		loggedInUser, session, err = c.authService.Login(context.Background(), username, password, code)
	}
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
	mfaStatus := "Disabled"
	if loggedInUser.TOTPEnabled {
		mfaStatus = "Enabled"
	}
	fmt.Println("MFA Status:", mfaStatus)
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

func (c *CLI) Exit() {
	fmt.Println("Goodbye.")
}

func (c *CLI) Clear() {
	fmt.Print("\033[H\033[2J")
}

func (c *CLI) Help() {
	fmt.Println("usage: /[command]")
	fmt.Println("commands:")
	fmt.Println("  /help        :Show help screen")
	fmt.Println("  /clear       :Clear the screen")
	fmt.Println("  /register    :Register user")
	fmt.Println("  /login       :Login with credentials")
	fmt.Println("  /exit        :Exit the app")
	fmt.Println("  /whoami      :Show current user")
	fmt.Println("  /logout      :Logout of current session")
	fmt.Println("  /enable2fa   :Enable TOTP two-factor authentication")
	fmt.Println("  /disable2fa  :Disable TOTP two-factor authentication")
	fmt.Println("")
	fmt.Println("Hit CRTL-C or CTRL-D to quit the application forcefully")
}

// after login
func (c *CLI) Whoami() {
	if !c.requireLogin() {
		return
	}
	fmt.Println("Current user:", c.currentUser.Username)
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

func (c *CLI) Enable2FA() {
	if !c.requireLogin() {
		return
	}

	secret, url, err := c.authService.Enable2FA(context.Background(), c.currentUser.ID, c.currentUser.Username)
	if err != nil {
		fmt.Println("\033[1;31mEnable 2FA failed:\033[0m", err)
		return
	}

	c.currentUser.TOTPEnabled = true
	c.currentUser.TOTPSecret = &secret
	fmt.Println("\033[1;32m2FA enabled successfully.\033[0m")
	fmt.Println("Add this account to your authenticator app using the following URL:")
	fmt.Println(url)
	fmt.Println("Secret:", secret)
}

func (c *CLI) Disable2FA() {
	if !c.requireLogin() {
		return
	}

	if err := c.authService.Disable2FA(context.Background(), c.currentUser.ID); err != nil {
		fmt.Println("\033[1;31mDisable 2FA failed:\033[0m", err)
		return
	}

	c.currentUser.TOTPEnabled = false
	c.currentUser.TOTPSecret = nil
	fmt.Println("\033[1;32m2FA disabled successfully.\033[0m")
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
