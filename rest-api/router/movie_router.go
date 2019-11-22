package router

import (
	"github.com/gorilla/mux"

	. "poc-rest-api-go/rest-api/model"
	movieservice "poc-rest-api-go/rest-api/service"
)

type MovieRouter struct {
}

var router *mux.Router
var movieService *movieservice.MovieService

func (m *MovieRouter) Initialize(routerConfig *RouterConfig) {
	router = routerConfig.Router
	movieService = movieservice.NewMovieService(routerConfig.Database)

	router.HandleFunc("/movies", movieService.GetAll).Methods("GET")
	router.HandleFunc("/movies/{id}", movieService.GetById).Methods("GET")
	router.HandleFunc("/movies", movieService.Create).Methods("POST")
	router.HandleFunc("/movies/{id}", movieService.Update).Methods("PUT")
	router.HandleFunc("/movies/{id}", movieService.Delete).Methods("DELETE")
}
