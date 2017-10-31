package util

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

// onnectPG Ceates a connection to a postgresql DB, exits if there is an error
func ConnectPG(config PGConfig) *sql.DB {
	db, err := ConnectDBErr(config)
	CheckErrFatal(err)
	return db
}

// ConnectPGErr Creates a connection to a postgres DB, returns an error if failed
func ConnectPGErr(config PGConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host, config.User, config.Password, config.DB, config.Port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		LogErr(err)
		return db, err
	}
	err = db.Ping()
	if err != nil {
		LogErr(err)
		return db, err
	}
	return db, nil
}

// BytesToIntSlice turns a slice of bytes to a slice of ints of a user specified size
func BytesToIntSlice(bytes []byte, intBits int) ([]int64, error) {
	baseStr := PgArrayToString(bytes)
	intSlice := make([]int64, 0)
	for _, memberStr := range strings.Split(baseStr, ",") {
		member, err := strconv.ParseInt(memberStr, 10, intBits)
		if err != nil {
			return intSlice, err
		}
		intSlice = append(intSlice, member)
	}
	return intSlice, nil
}

// BytesToFloatSlice Turnas slice of bytes to a slice of ints of a users specified size
func BytesToFloatSlice(bytes []byte, floatBits int) ([]float64, error) {
	floatSlice := make([]float64, 0)
	for _, memberStr := range strings.Split(PgArrayToString(bytes), ",") {
		member, err := strconv.ParseFloat(memberStr, floatBits)
		if err != nil {
			return floatSlice, err
		}
		floatSlice = append(floatSlice, member)
	}
	return floatSlice, nil
}

// BytesToStrSlice turns a byte array into a string array
func BytesToStrSlice(bytes []byte) []string {
	baseStr := PgArrayToString(bytes)
	return strings.Split(baseStr, ",")
}

// PgArrayToString turns a byte arra into a string without enclosing {}
func PgArrayToString(bytes []byte) string {
	baseStr := string(bytes)
	return strings.Trim(baseStr, "{}")
}
