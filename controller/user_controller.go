package controller

import (
	"net/http"

	. "github.com/hongsongp97/tickethunter_server/dao"
	// "gopkg.in/mgo.v2/bson"
	// . "github.com/hongsongp97/tickethunter_server/model"
)

type UserController struct {
	dao     *DAO
	daoUser UserDAO
}

func (userController *UserController) Init() {
	userController.dao = mainDao
}

func (userController *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	// respondWithJson(w, http.StatusOK, "hehe")
}
