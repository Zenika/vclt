// certificateManager
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/env/envHelpers.go
// Original timestamp: 2023/08/19 10:02

package env

import (
	"encoding/json"
	cerr "github.com/jeanfrancoisgratton/customError"
	"os"
	"path/filepath"
	"strings"
)

var EnvCfgFile string

type EnvironmentStruct struct {
	VaultAddress  string `json:"VaultAddress"`
	VaultUsername string `json:"VaultUsername"`
	VaultPassword string `json:"VaultPassword,omitempty"`
	KVpath        string `json:"KVpath"`
}

// Load the JSON env file in the user's .config/certificatemanager directory, and store it into a data type (struct)
func (e *EnvironmentStruct) LoadEnvironmentFile() *cerr.CustomError {
	var payload EnvironmentStruct
	var err error

	if !strings.HasSuffix(EnvCfgFile, ".json") {
		EnvCfgFile += ".json"
	}
	rcFile := filepath.Join(os.Getenv("HOME"), ".config", "JFG", "vclt", EnvCfgFile)
	jFile, err := os.ReadFile(rcFile)
	if err != nil {
		return &cerr.CustomError{Title: "Error reading config file", Message: err.Error()}
	}
	err = json.Unmarshal(jFile, &payload)
	if err != nil {
		return &cerr.CustomError{Title: err.Error(), Fatality: cerr.Fatal}
	}
	return nil
}

// Save the above structure into a JSON file in the user's .config/certificatemanager directory
func (e EnvironmentStruct) SaveEnvironmentFile(outputfile string) *cerr.CustomError {
	if outputfile == "" {
		outputfile = EnvCfgFile
	}
	jStream, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return &cerr.CustomError{Title: err.Error(), Fatality: cerr.Fatal}
	}
	rcFile := filepath.Join(os.Getenv("HOME"), ".config", "JFG", "certificatemanager", outputfile)
	if err = os.WriteFile(rcFile, jStream, 0600); err != nil {
		return &cerr.CustomError{Title: err.Error(), Fatality: cerr.Fatal}
	}

	return nil
}
