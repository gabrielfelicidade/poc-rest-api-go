package model

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type RouterConfig struct {
	Database *mongo.Database
	Router   *mux.Router
}
