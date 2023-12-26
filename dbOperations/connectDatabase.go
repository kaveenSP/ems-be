package dbOperations

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const DB_URI = ""
const DB_NAME = "Planners"

var DB *mongo.Database

func ConnectMongoDatabase() {
	dbClient, err := mongo.NewClient(options.Client().ApplyURI(DB_URI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err = dbClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//set database name
	DB = dbClient.Database(DB_NAME)
}
