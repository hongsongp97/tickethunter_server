package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/hongsongp97/tickethunter_server/config"
	. "github.com/hongsongp97/tickethunter_server/dao"
	"github.com/hongsongp97/tickethunter_server/router"
)

var config = Config{}

// Parse the configuration file 'config.toml'.
func init() {
	// Log with line number
	log.SetFlags(log.Lshortfile)

	// Read configuration file
	config.Read()
}

// Define HTTP request routes.
func main() {
	route := mux.NewRouter()

	// Establish a connection to DB.
	var dao = Dao{
		Server:   config.Server,
		Database: config.Database}
	dao.ConnectToDB()
	defer dao.Disconnect()

	// Set router
	router.SetUserRouter(route, &dao)
	router.SetEventRouter(route, &dao)

	if err := http.ListenAndServe(config.Host+config.Port, route); err != nil {
		log.Fatal(err)
	}
}
