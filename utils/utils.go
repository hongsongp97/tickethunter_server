package utils

import (
	"encoding/json"
	"net/http"

	. "github.com/hongsongp97/tickethunter_server/model"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	response := ResponseJson{Status: code, Message: msg}
	json.NewEncoder(w).Encode(response)
	// respondWithJson(w, code, map[string]string{"error": msg})
}

func RespondWithJson(w http.ResponseWriter, code int, data interface{}) {
	response := ResponseJson{Status: code, Data: data}
	json.NewEncoder(w).Encode(response)
	// response, _ := json.Marshal(payload)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(code)
	// w.Write(response)
}
