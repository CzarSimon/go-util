package main

import "log"

//CheckErr logs an error if present
func CheckErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

//CheckErrFatal logs an error if present and if so also exits
func CheckErrFatal(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
