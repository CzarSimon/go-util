package util

import (
	"database/sql"
	"fmt"
	"log"
)

//CheckErr logs an error if present
func CheckErr(err error) {
	if IsErr(err) {
		LogErr(err)
	}
}

//CheckErrFatal logs an error if present and if so also exits
func CheckErrFatal(err error) {
	if IsErr(err) {
		log.Fatalln(err.Error())
	}
}

//CheckErrAndRollback check and error and if present logs the error and
//performs a rollback on the supplied transactoin
func CheckErrAndRollback(err error, tx *sql.Tx) bool {
	if IsErr(err) {
		LogErr(err)
		tx.Rollback()
		return true
	}
	return false
}

// LogErr logs the supplied error to the console
func LogErr(err error) {
	log.Println(err.Error())
}

// PrintErr Prints the supplied error to stdout
func PrintErr(err error) {
	fmt.Println(err.Error())
}

//IsErr checks if an error is not nil
func IsErr(err error) bool {
	return err != nil
}
