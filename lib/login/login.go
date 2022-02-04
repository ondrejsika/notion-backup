package login

import (
	"fmt"

	"github.com/ondrejsika/notion-backup/lib/client"
)

func Login() error {
	var email string
	var password string

	api := client.New("")
	fmt.Println("Enter Your Email:")
	fmt.Scanln(&email)

	state, csrf, _ := api.SendTemporaryPassword(email)

	fmt.Println("Enter Your Password:")
	fmt.Scanln(&password)

	token, _ := api.LoginWithEmail(state, csrf, password)
	fmt.Println("--")
	fmt.Println(token)

	return nil
}
