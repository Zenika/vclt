// vclt
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : ls.go
// Original timestamp : 2024/06/30 16:39

package kv

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
	"vclt/sys"
)

func ListFields(entry string, version int) (map[string]interface{}, *cerr.CustomError) {
	var err error
	client, kvPath, ce := sys.Login(false)
	isVersioned, ce := sys.IsKVv2(client, kvPath)
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

	return data, nil
}

func ListEntries() ([]string, *cerr.CustomError) {
	var err error
	client, kvPath, ce := sys.Login(false)
	if ce != nil {
		return nil, ce
	}
	isVersioned, ce := sys.IsKVv2(client, kvPath)
	if ce != nil {
		return nil, ce
	}

	var secret *vault.Secret
	if isVersioned {
		// KV v2 path
		secret, err = client.Logical().List(fmt.Sprintf("%s/metadata", kvPath))
	} else {
		// KV v1 path
		secret, err = client.Logical().List(kvPath)
	}

	if err != nil {
		return nil, &cerr.CustomError{Title: "failed to read from KV store", Message: err.Error()}
	}

	if secret == nil || secret.Data == nil {
		return nil, &cerr.CustomError{Title: fmt.Sprintf("no data found at %s", kvPath)}
	}

	entries, ok := secret.Data["keys"].([]interface{})
	if !ok {
		return nil, &cerr.CustomError{Title: "failed to parse entries"}
	}

	var result []string
	for _, entry := range entries {
		result = append(result, entry.(string))
	}

	return result, nil
}
