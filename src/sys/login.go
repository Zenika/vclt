// vclt
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : login.go
// Original timestamp : 2024/06/28 15:38

package sys

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"vclt/env"
)

var Quiet = false

func Login(showRes bool) (*vault.Client, string, *cerr.CustomError) {
	var e env.EnvironmentStruct
	var ce *cerr.CustomError

	config := vault.DefaultConfig()
	if e, ce = env.LoadEnvironmentFile(); ce != nil {
		return nil, "", ce
	}

	options := map[string]interface{}{
		"password": hf.DecodeString(e.VaultPassword, ""),
	}

	path := fmt.Sprintf("auth/userpass/login/%s", e.VaultUsername)
	config.Address = e.VaultAddress
	client, err := vault.NewClient(config)
	if err != nil {
		return nil, "", &cerr.CustomError{Title: "Failed to create Vault client:", Message: err.Error()}
	}

	secret, err := client.Logical().Write(path, options)
	if err != nil {
		return nil, "", &cerr.CustomError{Title: "Failed to login", Message: err.Error()}
	}

	client.SetToken(secret.Auth.ClientToken)
	if showRes {
		fmt.Printf("Login %s\n", hf.Green("succeeded"))
	}
	return client, e.KeyValuePath, nil
}
