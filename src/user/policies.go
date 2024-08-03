package user

import (
	vault "github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
	"vclt/sys"
)

func AssignPolicyToUser(client *vault.Client, username string, policies []string) *cerr.CustomError {
	client, _, ce := sys.Login(false)
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
