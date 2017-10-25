package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// StatusOK The status 200 return message
const StatusOK string = "200 - OK"

// StatusUnauthorized The status 401 return message
const StatusUnauthorized string = "401 - Unauthorized"

//SendPlainTextRes sends a plain text message given as an input
func SendPlainTextRes(res http.ResponseWriter, msg string) {
	res.Header().Set("Content-Type", "text/plain")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(msg))
}

//JSONRes contains a response
type JSONRes struct {
	Response string
}

//SendJSONStringRes serializes a given message sting to JSON and sends to the requestor
func SendJSONStringRes(res http.ResponseWriter, msg string) {
	jsonResponse := JSONRes{msg}
	js, err := json.Marshal(jsonResponse)
	if err != nil {
		SendErrRes(res, err)
		return
	}
	SendJSONRes(res, js)
}

//SendJSONRes send a given JSON respons to the requestor
func SendJSONRes(res http.ResponseWriter, js []byte) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.WriteHeader(http.StatusOK)
	res.Write(js)
}

//SendErrRes sends a given error to the requestor
func SendErrRes(res http.ResponseWriter, err error) {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

// SendErrStatus Sends an error and a given status code to the requestor
func SendErrStatus(res http.ResponseWriter, err error, statusCode int) {
	if err != nil {
		http.Error(res, err.Error(), statusCode)
	}
}

//SendOK sends status code 200 to the requestor
func SendOK(res http.ResponseWriter) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(StatusOK))
}

// Ping sends a 200 response if the server is running
func Ping(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Ping recieved")
	SendOK(res)
}

// SendUnauthorized Sends an unauthorized status to the requestor
func SendUnauthorized(res http.ResponseWriter) {
	http.Error(res, StatusUnauthorized, http.StatusUnauthorized)
}

//PlaceholderHandler is a dummy handler to ease development
func PlaceholderHandler(res http.ResponseWriter, req *http.Request) {
	SendOK(res)
}

// ParseValueFromQuery Parses a query value from request
func ParseValueFromQuery(req *http.Request, key, errorMsg string) (string, error) {
	value := req.URL.Query().Get(key)
	if value == "" {
		return "", errors.New(errorMsg)
	}
	return value, nil
}
