package utils

import (
	"encoding/json"
	"log"
	"net/http"

	. "github.com/hongsongp97/tickethunter_server/models"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	response := ResponseJson{Status: code, Message: msg}
	prettyRes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Println(err)
	}
	w.Write(prettyRes)
}

func RespondWithJson(w http.ResponseWriter, code int, data interface{}) {
	response := ResponseJson{Status: code, Data: data}
	prettyRes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Println(err)
	}
	w.Write(prettyRes)
}
