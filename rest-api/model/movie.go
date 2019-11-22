package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Active      bool               `json:"active" bson:"active"`
}
