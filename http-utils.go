package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//SendPlainTextRes sends a plain text message given as an input
func SendPlainTextRes(res http.ResponseWriter, msg string) {
	res.Header().Set("Content-Type", "text/plain")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Write([]byte(fmt.Sprintf("%x", msg)))
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
	res.Write(js)
}

//SendErrRes sends a given error to the requestor
func SendErrRes(res http.ResponseWriter, err error) {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

//SendOK sends status code 200 to the requestor
func SendOK(res http.ResponseWriter) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("200 - OK"))
}

//PlaceholderHandler is a dummy handler to ease development
func PlaceholderHandler(res http.ResponseWriter, req *http.Request) {
	SendOK(res)
}
