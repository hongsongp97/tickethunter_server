package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/hongsongp97/tickethunter_server/config"
	. "github.com/hongsongp97/tickethunter_server/controllers"
	. "github.com/hongsongp97/tickethunter_server/dao"
)

var config = Config{}
var dao = Dao{}
var userController = UserController{}

// Parse the configuration file 'config.toml', and establish a connection to DB
// Init controllers for routes
func init() {
	// Log with line number
	log.SetFlags(log.Lshortfile)

	// Read configuration file
	config.Read()

	// Establish a connection to DB
	dao.Server = config.Server
	dao.Database = config.Database
	dao.ConnectToDB()

	// Init controller for user route
	userController.Init(&dao)
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user", userController.GetAllUsers).Methods("GET")
	r.HandleFunc("/user", userController.AddUser).Methods("POST")
	r.HandleFunc("/user", userController.UpdateUser).Methods("PUT")
	r.HandleFunc("/user", userController.DeleteUserById).Methods("DELETE")
	r.HandleFunc("/user/{id}", userController.GetUserById).Methods("GET")

	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatal(err)
	}
}
