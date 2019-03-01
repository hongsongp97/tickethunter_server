package controller

import (
	"encoding/json"
	"net/http"

	. "github.com/hongsongp97/tickethunter_server/dao"
	. "github.com/hongsongp97/tickethunter_server/model"
	"github.com/hongsongp97/tickethunter_server/utils"
)

var userDAO UserDAO
var dao DAO

var getAllUsers = func(w http.ResponseWriter, r *http.Request, dao *DAO) {

	userDAO.Dao = dao

	var users []User
	users, err := userDAO.FindAll()
	if err != nil {
		// json.NewEncoder(w).Encode()
		&utils.RespondWithJson(w, http.StatusNotFound, users)
		return
	}
	json.NewEncoder(w).Encode(users)

}
