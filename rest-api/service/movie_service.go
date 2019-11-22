package service

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

	. "poc-rest-api-go/rest-api/model"
	moviesrepository "poc-rest-api-go/rest-api/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type MovieService struct {
	MoviesRepository *moviesrepository.MoviesRepository
}

func NewMovieService(db *mongo.Database) *MovieService {
	return &MovieService{moviesrepository.NewMoviesRepository(db)}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (m *MovieService) GetAll(w http.ResponseWriter, r *http.Request) {
	movies, err := m.MoviesRepository.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

func (m *MovieService) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	movie, err := m.MoviesRepository.GetByID(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}

func (m *MovieService) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = primitive.NewObjectID()
	if err := m.MoviesRepository.Create(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, movie)
}

func (m *MovieService) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	id, _ := primitive.ObjectIDFromHex(params["id"])
	movie.ID = id
	if err := m.MoviesRepository.Update(id, movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": movie.Name + " atualizado com sucesso!"})
}

func (m *MovieService) Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	if err := m.MoviesRepository.Delete(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
