package util

import (
	"encoding/json"
	"log"
	"net/http"

	. "github.com/hongsongp97/tickethunter_server/model"
)

// RespondWithError func to send error response with pretty-json format
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	response := ResponseJson{Status: code, Message: msg}
	prettyRes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Println(err)
	}
	w.Write(prettyRes)
}

// RespondWithError func to send success response with pretty-json format
func RespondWithJson(w http.ResponseWriter, code int, data interface{}) {
	response := ResponseJson{Status: code, Data: data}
	prettyRes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Println(err)
	}
	w.Write(prettyRes)
}
