package env

import (
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
)

func CreateEnvFile(fname string) *cerr.CustomError {
	if EnvName == "" || VAddress == "" || VUserName == "" || VPassword == "" || KVstorePath == "" {
		return &cerr.CustomError{Fatality: cerr.Warning, Title: "Missing parameters", Message: "Use : vclt env create -h for more info", Code: 2}
	}

	es := EnvironmentStruct{
		EnvironmentName: EnvName,
		VaultAddress:    VAddress,
		VaultUsername:   VUserName,
		VaultPassword:   hf.EncodeString(VPassword, ""),
		KeyValuePath:    KVstorePath,
		Comments:        EnvComments,
	}

	return es.SaveEnvironmentFile(fname)
}
