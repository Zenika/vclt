// vclt
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : user.go
// Original timestamp : 2024/06/30 18:16

package sys

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
)

func CreateUser(username string, password string) *cerr.CustomError {
	client, _, ce := Login(false)
	if ce != nil {
		return ce
	}
	path := "auth/userpass/users/" + username
	if password == "" {
		password = hf.GetPassword(fmt.Sprintf("Please provide a password for %s", username))
	}
	data := map[string]interface{}{
		"password": password,
	}

	_, err := client.Logical().Write(path, data)
	if err != nil {
		return &cerr.CustomError{Title: err.Error()}
	}
	return nil
}

func AssignPolicyToUser(client *vault.Client, username string, policies []string) *cerr.CustomError {
	client, _, ce := Login(false)
	if ce != nil {
		return ce
	}
	path := "auth/userpass/users/" + username
	data := map[string]interface{}{
		"policies": policies,
	}

	_, err := client.Logical().Write(path, data)
	if err != nil {
		return &cerr.CustomError{Title: err.Error()}
	}
	return nil
}
