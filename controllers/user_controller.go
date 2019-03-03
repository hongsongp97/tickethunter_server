/*
controllers is packages
*/
package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/hongsongp97/tickethunter_server/dao"
	. "github.com/hongsongp97/tickethunter_server/models"
	"github.com/hongsongp97/tickethunter_server/utils"
)

type UserController struct {
	UserDao UserDao
}

func (userController *UserController) Init(dao *Dao) {
	userController.UserDao = UserDao{Dao: dao}
	userController.UserDao.Init()

	log.SetFlags(log.Lshortfile)
}

func HandlingResponseErrorWithJson(err error) {

}

func (userController *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	users, err := userController.UserDao.FindAll()
	if err != nil {
		utils.RespondWithJson(w, http.StatusBadRequest, "Cannot get data from database.")
		return
	}
	if len(users) == 0 {
		utils.RespondWithJson(w, http.StatusBadRequest, "Empty data.")
	}
	utils.RespondWithJson(w, http.StatusOK, users)
}

func (userController *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	var user User
	params := mux.Vars(r)
	user, err := userController.UserDao.FindById(string(params["id"]))

	if err != nil {
		utils.Re
	}

}
