package moviesrepository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	. "poc-rest-api-go/rest-api/model"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	COLLECTION_NAME = "movies"
)

type MoviesRepository struct {
	Database   *mongo.Database
	Collection *mongo.Collection
}

func NewMoviesRepository(db *mongo.Database) *MoviesRepository {
	return &MoviesRepository{db, db.Collection(COLLECTION_NAME)}
}

func (m *MoviesRepository) GetAll() ([]Movie, error) {
	var movies []Movie
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := m.Collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result Movie
		err := cur.Decode(&result)
		if err != nil {
			log.Println(err)
			return nil, nil
		}
		movies = append(movies, result)
	}
	return movies, err
}

func (m *MoviesRepository) GetByID(id primitive.ObjectID) (Movie, error) {
	var result Movie
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := m.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func (m *MoviesRepository) Create(movie Movie) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := m.Collection.InsertOne(ctx, movie)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *MoviesRepository) Update(id primitive.ObjectID, movie Movie) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := m.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": movie})
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *MoviesRepository) Delete(id primitive.ObjectID) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := m.Collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Println(err)
	}
	return err
}
