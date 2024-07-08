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

var ConfigFile string
var EnvName, VAddress, VUserName, VPassword, KVstorePath, EnvComments string

type EnvironmentStruct struct {
	EnvironmentName string `json:"EnvironmentName"`
	VaultAddress    string `json:"VaultAddress"`
	VaultUsername   string `json:"VaultUsername"`
	VaultPassword   string `json:"VaultPassword,omitempty"`
	KeyValuePath    string `json:"KeyValuePath"`
	Comments        string `json:"Comments,omitempty"`
}

// LoadEnvironmentFile : Load the JSON env file in the user's .config/JFG/vclt directory, and store it into a data type (struct)
func LoadEnvironmentFile() (EnvironmentStruct, *cerr.CustomError) {
	var payload EnvironmentStruct

	var ce *cerr.CustomError

	if !strings.HasSuffix(ConfigFile, ".json") {
		ConfigFile += ".json"
	}

	rcFile := filepath.Join(os.Getenv("HOME"), ".config", "JFG", "vclt", ConfigFile)
	_, err := os.Stat(rcFile)
	if os.IsNotExist(err) {
		payload, ce = getEnvParams(rcFile)
		if ce != nil {
			return EnvironmentStruct{}, ce
		}
		es := &payload
		if ce = es.SaveEnvironmentFile(ConfigFile); ce != nil {
			return EnvironmentStruct{}, ce
		}

	}

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

// SaveEnvironmentFile : Save the above structure into a JSON file in the user's .config/JFG/vclt directory
func (e *EnvironmentStruct) SaveEnvironmentFile(outputfile string) *cerr.CustomError {
	if outputfile == "" {
		outputfile = ConfigFile
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

// getEnvParams : Prompts for values to fill up the Environment structure
func getEnvParams(cfgFile string) (EnvironmentStruct, *cerr.CustomError) {
	es := EnvironmentStruct{}

	es.EnvironmentName = hf.GetStringValFromPrompt("Enter the name of this environment : ")
	es.VaultAddress = hf.GetStringValFromPrompt("Enter the address of the Vault (ex: https://mydomain:1234) : ")
	es.VaultUsername = hf.GetStringValFromPrompt("Please enter the username using the Vault : ")
	es.VaultPassword = hf.GetPassword("Please enter that user's password (leaving this empty will get you prompted for it, later) : ")
	if es.VaultPassword != "" {
		es.VaultPassword = hf.EncodeString(es.VaultPassword, "")
	}
	es.KeyValuePath = hf.GetStringValFromPrompt("Please enter the kv store path (no trailing/leading slash) : ")
	es.KeyValuePath = strings.TrimPrefix(es.KeyValuePath, "/")
	es.KeyValuePath = strings.TrimSuffix(es.KeyValuePath, "/")

	es.Comments = hf.GetStringValFromPrompt("(OPTIONAL) Please enter a comment : ")
	// Call the SaveEnvironmentFile() method using a pointer to es
	if err := es.SaveEnvironmentFile(cfgFile); err != nil {
		return EnvironmentStruct{}, err
	}

	return es, nil
}
