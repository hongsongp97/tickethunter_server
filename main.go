package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/hongsongp97/tickethunter_server/config"
	. "github.com/hongsongp97/tickethunter_server/dao" // "gopkg.in/mgo.v2/bson"
	. "github.com/hongsongp97/tickethunter_server/model"
)

var config = Config{}
var dao = DAO{}
var userDAO = UserDAO{}

// GET list of users
func AllUserEndPoint(w http.ResponseWriter, r *http.Request) {
	// users, err := userDAO.FindAll()
	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// user := User{FirstName: "Daniel", LastName: "Pham"}
	// respondWithJson(w, http.StatusOK, user)
	respondWithError(w, http.StatusInternalServerError, "Null response")
}

// GET a user by its ID
// func FindUserEndpoint(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	user, err := userDAO.FindById(params["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
// 		return
// 	}
// 	respondWithJson(w, http.StatusOK, user)
// }

// POST a new user
// func CreateUserEndPoint(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	var user User
// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	user.ID = bson.NewObjectId()
// 	if err := userDAO.Insert(user); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJson(w, http.StatusCreated, user)
// }

// // PUT update an existing user
// func UpdateUserEndPoint(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	var user User
// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	if err := dao.Update(user); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
// }

// // DELETE an existing user
// func DeleteuserEndPoint(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	var user User
// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	if err := dao.Delete(user); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
// }

func respondWithError(w http.ResponseWriter, code int, msg string) {
	response := ResponseJson{Status: code, Message: msg}
	json.NewEncoder(w).Encode(response)
	// respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, data interface{}) {
	response := ResponseJson{Status: code, Data: data}
	json.NewEncoder(w).Encode(response)
	// response, _ := json.Marshal(payload)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(code)
	// w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.ConnectToDB()

	userDAO.Dao = &dao
	userDAO.Init()

	users, err := userDAO.FindAll()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(users)
	}
	// userDAO.Insert()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", AllUserEndPoint).Methods("GET")
	// r.HandleFunc("/user", CreateUserEndPoint).Methods("POST")
	// r.HandleFunc("/user", UpdateUserEndPoint).Methods("PUT")
	// r.HandleFunc("/user", DeleteUserEndPoint).Methods("DELETE")
	// r.HandleFunc("/user/{id}", FindUserEndpoint).Methods("GET")
	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatal(err)
	}
}
