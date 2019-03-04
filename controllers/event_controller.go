package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/hongsongp97/tickethunter_server/dao"
	. "github.com/hongsongp97/tickethunter_server/models"
	"github.com/hongsongp97/tickethunter_server/utils"
	. "github.com/hongsongp97/tickethunter_server/utils"
)

// EventController is type, its instance handle proccesses to use Data Access Object.
type EventController struct {
	eventDao EventDao
}

// Init() func initializes UserDao of EventController, to access DB.
// Must call this func before using any other methods.
func (eventController *EventController) Init(dao *Dao) {
	// var eventDao = EventDao{Dao: dao}
	// eventController.eventDao = eventDao
	eventController.eventDao.Init(dao)
}

//GetEventById is used to get Event by ID
func (eventController *EventController) GetEventById(w http.ResponseWriter, r *http.Request) {
	var (
		params = mux.Vars(r)
		event  Event
		err    error
	)

	event, err = eventController.eventDao.FindById(params["id"])

	// handle with error
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Cannot get data")
		return
	}
	RespondWithJson(w, http.StatusOK, event)
}

func (eventController *EventController) GetEvents(w http.ResponseWriter, r *http.Request) {
	var (
		events []Event
		err    error
	)
	events, err = eventController.eventDao.FindAll()

	switch {
	case err != nil:
		RespondWithError(w, http.StatusBadRequest, "Cannot get data")

	case len(events) == 0:
		RespondWithJson(w, http.StatusBadRequest, "Empty data")
	default:
		RespondWithJson(w, http.StatusOK, events)
	}

	return

}

func (eventController *EventController) CreateEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var event Event

	err := json.NewDecoder(r.Body).Decode(&event)
	switch {
	case err != nil:
		RespondWithError(w, http.StatusBadRequest, "Cannot create new event! Invaild input!")
		log.Println("Create Event Error:" + err.Error())
		return
	case event.Id == "":
		event.Id = bson.NewObjectId().Hex()
	}
	if err := eventController.eventDao.Insert(event); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Cannot add new event!")
		log.Println("Create Event Error:" + err.Error())
		return
	}

	RespondWithJson(w, http.StatusOK, "add new event succesfully!")
}

func (eventController *EventController) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var event Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Cannot update event! Invaild input!")
		log.Println("Update Event Error:" + err.Error())
		return
	}

	if err := eventController.eventDao.Update(event); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Cannot update event!")
		log.Println("Update Event Error:" + err.Error())
		return
	}

	RespondWithJson(w, http.StatusOK, "Updated event succesfully!")

}

func (eventController *EventController) DeleteEvent(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var params = mux.Vars(r)
	if err := eventController.eventDao.Delete(params["id"]); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot delete event!")
		log.Println("Delete Event Error:" + err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, "Deleted event successfuly!")
}

func (eventController *EventController) GetJoinedUserByEventId(w http.ResponseWriter, r *http.Request) {
	var (
		params          = mux.Vars(r)
		userByEachEvent []UserByEachEvent
	)

	userByEachEvent, err := eventController.eventDao.FindByJoinedUser(params["id"])

	switch {
	case err != nil:
		RespondWithError(w, http.StatusBadRequest, "Cannot get data")

	case len(userByEachEvent) == 0:
		RespondWithJson(w, http.StatusBadRequest, "Empty data")
	default:
		RespondWithJson(w, http.StatusOK, userByEachEvent)
	}

	return

}

func (eventController *EventController) GetFollowedUserByEventId(w http.ResponseWriter, r *http.Request) {
	var (
		params          = mux.Vars(r)
		userByEachEvent []UserByEachEvent
	)

	userByEachEvent, err := eventController.eventDao.FindByFollowedUser(params["id"])

	switch {
	case err != nil:
		RespondWithError(w, http.StatusBadRequest, "Cannot get data")

	case len(userByEachEvent) == 0:
		RespondWithJson(w, http.StatusBadRequest, "Empty data")
	default:
		RespondWithJson(w, http.StatusOK, userByEachEvent)
	}

	return

}
