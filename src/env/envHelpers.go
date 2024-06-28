// certificateManager
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/env/envHelpers.go
// Original timestamp: 2023/08/19 10:02

package env

import (
	"encoding/json"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"os"
	"path/filepath"
	"strings"
)

var EnvCfgFile string

type EnvironmentStruct struct {
	VaultAddress  string `json:"VaultAddress"`
	VaultUsername string `json:"VaultUsername"`
	VaultPassword string `json:"VaultPassword,omitempty"`
	KeyValuePath  string `json:"KeyValuePath"`
}

// Load the JSON env file in the user's .config/JFG/vclt directory, and store it into a data type (struct)
func LoadEnvironmentFile() (EnvironmentStruct, *cerr.CustomError) {
	var payload EnvironmentStruct
	var ce *cerr.CustomError

	if !strings.HasSuffix(EnvCfgFile, ".json") {
		EnvCfgFile += ".json"
	}

	_, err := os.Stat(EnvCfgFile)
	if os.IsNotExist(err) {
		payload, ce = createEnvCfgFile(EnvCfgFile)
		return payload, ce
	}

	rcFile := filepath.Join(os.Getenv("HOME"), ".config", "JFG", "vclt", EnvCfgFile)
	jFile, err := os.ReadFile(rcFile)
	if err != nil {
		return EnvironmentStruct{}, &cerr.CustomError{Title: "Error reading the file", Message: err.Error()}
	}
	err = json.Unmarshal(jFile, &payload)
	if err != nil {
		return EnvironmentStruct{}, &cerr.CustomError{Title: "Error unmarshalling JSON", Message: err.Error()}
	} else {
		return payload, nil
	}
}

// Save the above structure into a JSON file in the user's .config/JFG/vclt directory
func (e *EnvironmentStruct) SaveEnvironmentFile(outputfile string) *cerr.CustomError {
	if outputfile == "" {
		outputfile = EnvCfgFile
	}

	if !strings.HasSuffix(outputfile, ".json") {
		outputfile += ".json"
	}

	jStream, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return &cerr.CustomError{Title: "Error marshalling information", Message: err.Error()}
	}
	rcFile := filepath.Join(os.Getenv("HOME"), ".config", "JFG", "vclt", outputfile)
	if err = os.WriteFile(rcFile, jStream, 0600); err != nil {
		return &cerr.CustomError{Title: "Error writing config file", Message: err.Error()}
	}

	return nil
}

func createEnvCfgFile(cfgFile string) (EnvironmentStruct, *cerr.CustomError) {
	es := EnvironmentStruct{}

	es.VaultAddress = hf.GetStringValFromPrompt("Enter the address of the Vault (ex: https://mydomain:1234) : ")
	es.VaultUsername = hf.GetStringValFromPrompt("Please enter the username using the Vault : ")
	es.VaultPassword = hf.GetPassword("Please enter that user's password (leaving this empty will get you prompted for it, later) : ")
	if es.VaultPassword != "" {
		es.VaultPassword = hf.EncodeString(es.VaultPassword, "")
	}
	es.KeyValuePath = hf.GetStringValFromPrompt("Please enter the kv store path (no trailing/leading slash) : ")
	es.KeyValuePath = strings.TrimPrefix(es.KeyValuePath, "/")
	es.KeyValuePath = strings.TrimSuffix(es.KeyValuePath, "/")

	// Call the SaveEnvironmentFile() method using a pointer to es
	if err := es.SaveEnvironmentFile(cfgFile); err != nil {
		return EnvironmentStruct{}, err
	}

	return es, nil
}
