package util

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

//ConnectPG creates a connection to a postgresql DB
func ConnectPG(config PGConfig) *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.User, config.Password, config.DB)
	db, err := sql.Open("postgres", connStr)
	CheckErrFatal(err)
	err = db.Ping()
	CheckErrFatal(err)
	return db
}

//BytesToIntSlice turns a slice of bytes to a slice of ints of a user specified size
func BytesToIntSlice(bytes []byte, intBits int) ([]int64, error) {
	baseStr := PgArrayToString(bytes)
	intSlice := make([]int64, 0)
	for _, memberStr := range strings.Split(baseStr, ",") {
		member, err := strconv.ParseInt(memberStr, 10, intBits)
		if IsErr(err) {
			return intSlice, err
		}
		intSlice = append(intSlice, member)
	}
	return intSlice, nil
}

//BytesToStrSlice turns a byte array into a string array
func BytesToStrSlice(bytes []byte) []string {
	baseStr := PgArrayToString(bytes)
	return strings.Split(baseStr, ",")
}

//PgArrayToString turns a byte arra into a string without enclosing {}
func PgArrayToString(bytes []byte) string {
	baseStr := string(bytes)
	return strings.Trim(baseStr, "{}")
}
