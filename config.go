package util

import (
	"fmt"
	"os"
)

//PGConfig holds connection parameters for a postgresql db.
type PGConfig struct {
	Host, Port, Password, User, DB string
}

//GetPGConfig creates a PGConfig based on environment variables and default values
func GetPGConfig(stdHost, stdPwd, stdUser, stdDB string) PGConfig {
	return PGConfig{
		Host:     GetEnvVar("DB_HOST", stdHost),
		Port:     GetEnvVar("DB_PORT", "5432"),
		Password: GetEnvVar("DB_PASSWORD", stdPwd),
		User:     GetEnvVar("DB_USER", stdUser),
		DB:       GetEnvVar("DB_NAME", stdDB),
	}
}

//ServerConfig contains setup and connection info for a server
type ServerConfig struct {
	Host, Port, Protocol string
}

//ToURL accepts a route as a parameter and turns the route + ServerConfig to a full url
func (server ServerConfig) ToURL(route string) string {
	return fmt.Sprintf("%s://%s:%s/%s", server.Protocol, server.Host, server.Port, route)
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
