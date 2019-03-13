package router

import (
	. "github.com/gorilla/mux"
	. "github.com/hongsongp97/tickethunter_server/controller"
	. "github.com/hongsongp97/tickethunter_server/dao"
)

var userController = UserController{}

func SetUserRouter(r *Router, dao *Dao) {
	// Init controller for user route
	userController.Init(dao)

	r.HandleFunc("/user", userController.GetAllUsers).Methods("GET")
	r.HandleFunc("/user", userController.AddUser).Methods("POST")
	r.HandleFunc("/user", userController.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", userController.DeleteUserById).Methods("DELETE")
	r.HandleFunc("/user/{id}", userController.GetUserById).Methods("GET")
}
