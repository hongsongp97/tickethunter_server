package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/hongsongp97/tickethunter_server/dao"
	. "github.com/hongsongp97/tickethunter_server/models"
	. "github.com/hongsongp97/tickethunter_server/utils"
)

// EventController is type, its instance handle proccesses to use Data Access Object.
type EventController struct {
	eventDao EventDao
}

// Init() func initializes UserDao of EventController, to access DB.
// Must call this func before using any other methods.
func (eventController *EventController) Init(dao *Dao) {
	var eventDao = EventDao{Dao: dao}
	eventController.eventDao = eventDao
	eventController.eventDao.Init()
}

func (eventController *EventController) GetEventById(w http.ResponseWriter, r *http.Request) {
	var (
		params = mux.Vars(r)
		events Event
		err    error
	)

	events, err = eventController.eventDao.FindById(params["id"])

	// handle with error
	if err != nil {
		RespondWithJson(w, http.StatusBadRequest, "Cannot get data")
		return
	}
	RespondWithJson(w, http.StatusOK, events)
}
