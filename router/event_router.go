package router

import (
	. "github.com/gorilla/mux"
	. "github.com/hongsongp97/tickethunter_server/controller"
	. "github.com/hongsongp97/tickethunter_server/dao"
)

var eventController = EventController{}

func SetEventRouter(r *Router, dao *Dao) {
	eventController.Init(dao)

	r.HandleFunc("/event/{id}", eventController.GetEventById).Methods("GET")
	r.HandleFunc("/event", eventController.GetEvents).Methods("GET")
	r.HandleFunc("/event", eventController.CreateEvent).Methods("POST")
	r.HandleFunc("/event", eventController.UpdateEvent).Methods("PUT")
	r.HandleFunc("/event/{id}", eventController.DeleteEvent).Methods("DELETE")
	r.HandleFunc("/user/joined/{id}", eventController.GetJoinedUserByEventId).Methods("GET")
	r.HandleFunc("/user/followed/{id}", eventController.GetFollowedUserByEventId).Methods("GET")
	r.HandleFunc("/event/category/{id}", eventController.GetEventByCategoryId).Methods("GET")
}
