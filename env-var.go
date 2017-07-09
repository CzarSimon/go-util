package util

import "os"

// EnvVar Key value pair for an enviornment variable
type EnvVar struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Set Sets environment varialbe
func (envVar EnvVar) Set() {
	os.Setenv(envVar.Key, envVar.Value)
}

//GetEnvVar gets an envrionment variable based on a key, returns a default value
//if the environment variable is not found
func GetEnvVar(varKey, nilValue string) string {
	envVar := os.Getenv(varKey)
	if envVar != "" {
		return envVar
	}
	return nilValue
}
