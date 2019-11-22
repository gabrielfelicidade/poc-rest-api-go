package main

import (
	"fmt"
	"log"
	"net/http"
	. "poc-rest-api-go/rest-api/config"
	. "poc-rest-api-go/rest-api/model"
	. "poc-rest-api-go/rest-api/router"

	"github.com/gorilla/mux"
)

var config = Config{}
var database = Database{}
var movieRouter = MovieRouter{}
var router *mux.Router
var routerConfig = RouterConfig{}

func main() {
	router := mux.NewRouter()

	config.Read()
	database.Connect(config)

	routerConfig.Router = router
	routerConfig.Database = database.Database

	movieRouter.Initialize(&routerConfig)

	fmt.Println("Server running in port:", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
