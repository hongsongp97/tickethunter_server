package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	. "github.com/hongsongp97/tickethunter_server/dao"
	. "github.com/hongsongp97/tickethunter_server/model"
	"github.com/hongsongp97/tickethunter_server/util"
	"gopkg.in/mgo.v2/bson"
)

// EventController is type, its instance handle proccesses to use Data Access Object.
type EventController struct {
	eventDao EventDao
}

// Init() func initializes UserDao of EventController, to access DB.
// Must call this func before using any other methods.
func (eventController *EventController) Init(dao *Dao) {
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
		util.RespondWithError(w, http.StatusBadRequest, "Cannot get data. \n"+err.Error())
		return
	}
	util.RespondWithJson(w, http.StatusOK, event)
}

//Get All Event
func (eventController *EventController) GetEvents(w http.ResponseWriter, r *http.Request) {
	var (
		events []Event
		err    error
	)
	events, err = eventController.eventDao.FindAll()

	switch {
	case err != nil:
		util.RespondWithError(w, http.StatusBadRequest, "Cannot get data")

	case len(events) == 0:
		util.RespondWithJson(w, http.StatusBadRequest, "Empty data")
	default:
		util.RespondWithJson(w, http.StatusOK, events)
	}

	return

}

//Create Event
func (eventController *EventController) CreateEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var event Event

	err := json.NewDecoder(r.Body).Decode(&event)
	switch {
	case err != nil:
		util.RespondWithError(w, http.StatusBadRequest, "Cannot create new event! Invaild input!")
		log.Println("Create Event Error:" + err.Error())
		return
	case event.Id == "":
		event.Id = bson.NewObjectId().Hex()
	}
	if err := eventController.eventDao.Insert(event); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Cannot add new event! \n"+err.Error())
		log.Println("Create Event Error:" + err.Error())
		return
	}

	util.RespondWithJson(w, http.StatusOK, "add new event succesfully!")
}

//Update Event
func (eventController *EventController) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var event Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Cannot update event! Invaild input! \n"+err.Error())
		log.Println("Update Event Error:" + err.Error())
		return
	}

	if err := eventController.eventDao.Update(event); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Cannot update event!")
		log.Println("Update Event Error:" + err.Error())
		return
	}

	util.RespondWithJson(w, http.StatusOK, "Updated event succesfully!")

}

// Delete Event
func (eventController *EventController) DeleteEvent(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var params = mux.Vars(r)
	if err := eventController.eventDao.Delete(params["id"]); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Cannot delete event! \n"+err.Error())
		log.Println("Delete Event Error:" + err.Error())
		return
	}
	util.RespondWithJson(w, http.StatusOK, "Deleted event successfuly!")
}

//Get list of User Id that already joined the Event
func (eventController *EventController) GetJoinedUserByEventId(w http.ResponseWriter, r *http.Request) {
	var (
		params          = mux.Vars(r)
		userByEachEvent []UserByEachEvent
	)

	userByEachEvent, err := eventController.eventDao.FindByJoinedUser(params["id"])

	switch {
	case err != nil:
		util.RespondWithError(w, http.StatusBadRequest, "Cannot get data. \n"+err.Error())

	case len(userByEachEvent) == 0:
		util.RespondWithJson(w, http.StatusBadRequest, "Empty data")
	default:
		util.RespondWithJson(w, http.StatusOK, userByEachEvent)
	}

	return
}

//Get list of User Id that already Followed the Event
func (eventController *EventController) GetFollowedUserByEventId(w http.ResponseWriter, r *http.Request) {
	var (
		params          = mux.Vars(r)
		userByEachEvent []UserByEachEvent
	)

	userByEachEvent, err := eventController.eventDao.FindByFollowedUser(params["id"])

	switch {
	case err != nil:
		util.RespondWithError(w, http.StatusBadRequest, "Cannot get data. \n"+err.Error())

	case len(userByEachEvent) == 0:
		util.RespondWithJson(w, http.StatusBadRequest, "Empty data")
	default:
		util.RespondWithJson(w, http.StatusOK, userByEachEvent)
	}

	return

}

// Get All Events that belong to the specify Category
func (eventController *EventController) GetEventByCategoryId(w http.ResponseWriter, r *http.Request) {

	var (
		params = mux.Vars(r)
		events []Event
	)
	log.Println("params GetEventByCategoryId :" + params["id"])
	events, err := eventController.eventDao.FindByCategoryId(strings.ToUpper(params["id"]))

	switch {
	case err != nil:
		util.RespondWithError(w, http.StatusBadRequest, "Cannot get data. \n"+err.Error())

	case len(events) == 0:
		util.RespondWithJson(w, http.StatusBadRequest, "Empty data")
	default:
		util.RespondWithJson(w, http.StatusOK, events)
	}

	return
}
