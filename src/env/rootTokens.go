// vclt
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : rootTokens.go
// Original timestamp : 2024/08/03 12:09

package env

import (
	"encoding/json"
	"fmt"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"os"
	"path/filepath"
)

type rootKeysFile struct {
	ServerAddr string   `json:"ServerAddr"`
	Comment    string   `json:"Comment,omitempty"`
	Keypart    []string `json:"Keypart"`
}

func CreateRootKeys(mininmalNumberOfKeys int) *cerr.CustomError {
	rk := rootKeysFile{}

	fmt.Println("You will now be prompted to consign the root keys needed to unseal a vault")
	fmt.Println("You need at least 3 key parts to have a valid key. Press enter at the prompt to complete the process")
	fmt.Println("If you press ENTER before having entered 3 key parts, the whole process will be aborted")
	fmt.Printf("The keyparts will be stored in encrypted form in %s/.config/JFG/vclt/rootkeys.json\n", os.Getenv("HOME"))
	fmt.Println()

	fmt.Println(`You will now enter the server's address (URL).
If you simply press ENTER, we will use the VAULT_ADDRESS variable. If the variable is empty, we will abort`)
	rk.ServerAddr = hf.GetStringValFromPrompt("Please enter the Vault server URL: ")
	if rk.ServerAddr == "" {
		rk.ServerAddr = os.Getenv("VAULT_ADDRESS")
		if rk.ServerAddr == "" {
			return &cerr.CustomError{Fatality: cerr.Warning, Title: "VAULT_ADDRESS was not provided"}
		}
	}

	rk.Comment = hf.GetStringValFromPrompt("Enter a comment to describe this file, ENTER to skip")

	rk.Keypart = hf.GetStringSliceFromPrompt("Enter the root key parts you want to save on disk")
	if x := len(rk.Keypart); x < mininmalNumberOfKeys {
		return &cerr.CustomError{Fatality: cerr.Warning, Title: "Unmet the minimal number of key parts",
			Message: fmt.Sprintf("You only have %d parts out of the required %d.\n", x, mininmalNumberOfKeys)}
	}

	return writeJson(rk)
}

func writeJson(data rootKeysFile) *cerr.CustomError {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return &cerr.CustomError{Title: "Error marshalling information", Message: err.Error()}
	}

	// Write JSON data to file
	err = os.WriteFile(filepath.Join(os.Getenv("HOME"), ".config", "JFG", "vclt", "rootkeys.json"), jsonData, 0644)
	if err != nil {
		return &cerr.CustomError{Title: "Error writing config file", Message: err.Error()}
	}

	return nil
}
