package kv

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
)

//func Get(key string) string {
//	client, _ := Login()

func Get(entry string, field string, version int) (interface{}, *cerr.CustomError) {
	var err error
	client, kvPath, ce := Login(false)
	isVersioned, ce := IsKVv2(client, kvPath)
	if ce != nil {
		return nil, ce
	}
	var secret *vault.Secret

	if isVersioned {
		dataPath := fmt.Sprintf("%s/data/%s", kvPath, entry)
		queryParams := map[string][]string{}
		if version > 0 {
			queryParams["version"] = []string{fmt.Sprintf("%d", version)}
		}
		secret, err = client.Logical().ReadWithData(dataPath, queryParams)
	} else {
		secret, err = client.Logical().Read(entry)
	}

	if err != nil {
		return nil, &cerr.CustomError{Title: "failed to read from KV store", Message: err.Error()}
	}

	if secret == nil || secret.Data == nil {
		return nil, &cerr.CustomError{Title: fmt.Sprintf("no data found at %s", kvPath)}
	}

	var data map[string]interface{}
	if isVersioned {
		data = secret.Data["data"].(map[string]interface{})
	} else {
		data = secret.Data
	}

	if val, ok := data[field]; ok {
		return val, nil
	} else {
		return nil, &cerr.CustomError{Title: fmt.Sprintf("field %s not found in KV store", field)}
	}
}
