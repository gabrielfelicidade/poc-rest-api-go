package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Context  context.Context
	Database *mongo.Database
}

func (d *Database) Connect(c Config) {
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, cliErr := mongo.Connect(context, options.Client().ApplyURI(c.Server))
	if cliErr != nil {
		log.Fatal(cliErr)
	}
	d.Database = client.Database(c.Database)
}
