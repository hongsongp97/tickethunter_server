/*
controllers is package in charge of logical executing.
*/
package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/hongsongp97/tickethunter_server/dao"
	. "github.com/hongsongp97/tickethunter_server/models"
	"github.com/hongsongp97/tickethunter_server/utils"
	"gopkg.in/mgo.v2/bson"
)

// UserController is a type, its instance handle proccesses to use Data Access Object.
type UserController struct {
	UserDao UserDao //DAO for User from DB.
}

// Init() func initializes UserDao of UserController, to access DB.
// Must call this func before using any other methods.
func (userController *UserController) Init(dao *Dao) {
	log.SetFlags(log.Lshortfile)

	userController.UserDao = UserDao{Dao: dao}
	userController.UserDao.Init()
}

// GetAllUsers() is used to get all users record.
func (userController *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	users, err := userController.UserDao.FindAll()
	if err != nil {
		utils.RespondWithJson(w, http.StatusBadRequest, "Cannot get users data!")
		log.Println(err)
		return
	}
	if len(users) == 0 {
		utils.RespondWithJson(w, http.StatusBadRequest, "Empty data.")
	}
	utils.RespondWithJson(w, http.StatusOK, users)
}

// GetUserById() is used to get a user record by it's ID.
func (userController *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println(params["id"])
	user, err := userController.UserDao.FindById(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot get user data!")
		log.Println(err)
		return
	}
	utils.RespondWithJson(w, http.StatusOK, user)
}

// AddUser() is used to create a user record.
func (userController *UserController) AddUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot add new user, invalid input!")
		log.Println(err)
		return
	}
	if user.Id == "" {
		user.Id = bson.NewObjectId().Hex()
	}
	if err := userController.UserDao.Insert(user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot add new user!")
		log.Println(err)
		return
	}
	utils.RespondWithJson(w, http.StatusOK, "Add new user successfuly!")
}

// DeleteUserById() is used to delete a user record by it's ID.
func (userController *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var userId UserIdJson
	if err := json.NewDecoder(r.Body).Decode(&userId); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot delete user, invalid input!")
		log.Println(err)
		return
	}
	if err := userController.UserDao.Delete(userId.Value); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot delete user!")
		log.Println(err)
		return
	}
	utils.RespondWithJson(w, http.StatusOK, "Deleted user successfuly!")
}

// UpdateUser() is used to update a user record by it's ID.
func (userController *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot update user, invalid input!")
		log.Println(err)
		return
	}
	log.Println(user)
	if err := userController.UserDao.Update(user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot update user!")
		log.Println(err)
		return
	}
	utils.RespondWithJson(w, http.StatusOK, "Updated user successfuly!")
}
