// vclt
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : addRemoveUser.go
// Original timestamp : 2024/06/30 18:16

package user

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"vclt/sys"
)

// Create a new user. This assumes that the user of this app is properly logged in as
// Someone with sufficient rights to perform such task
func CreateUser(username string, password string) *cerr.CustomError {
	client, _, ce := sys.Login(false)
	if ce != nil {
		return ce
	}
	path := "auth/userpass/users/" + username
	if password == "" {
		password = hf.GetPassword(fmt.Sprintf("Please provide a password for %s", username))
	}
	return writeUserData(client, path, password)
}

// Change a given user's password
func ChangePassword(username string, password string) *cerr.CustomError {
	client, _, ce := sys.Login(false)
	if ce != nil {
		return ce
	}

	path := fmt.Sprintf("auth/userpass/users/%s/password", username)

	return writeUserData(client, path, password)
}

// This is a common code path used by both CreateUser() and ChangePassword()
func writeUserData(cli *api.Client, path, password string) *cerr.CustomError {
	data := map[string]interface{}{
		"password": password,
	}

	_, err := cli.Logical().Write(path, data)
	if err != nil {
		return &cerr.CustomError{Title: err.Error()}
	}
	return nil
}
