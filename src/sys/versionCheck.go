// vclt
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : versionCheck.go
// Original timestamp : 2024/06/30 15:36

package sys

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
)

func IsKVv2(client *vault.Client, path string) (bool, *cerr.CustomError) {
	sys := client.Sys()
	mounts, err := sys.ListMounts()
	if err != nil {
		return false, &cerr.CustomError{Title: "failed to list mounts", Message: err.Error()}
	}

	for mountPath, mount := range mounts {
		if path+"/" == mountPath || path == mountPath {
			return mount.Options["version"] == "2", nil
		}
	}
	return false, &cerr.CustomError{Title: fmt.Sprintf("no mount found at path %s", path)}
}
