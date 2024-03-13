package db

import (
	"context"
	"log"
	"os"

	"github.com/vojtech-mikes/opentrail/internals/model"
	"github.com/vojtech-mikes/opentrail/internals/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: Change context type from TODO

func createDbConnection() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	dbUri := os.Getenv("DB_URI")

	if dbUri == "" {
		log.Panicf("Failed to get DB_URI")
	}

	opt := options.Client().ApplyURI(dbUri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opt)

	if err != nil {
		log.Panicf("Failed to create MongoDB client")
	}

	return client
}

func FindOne(filter interface{}) model.User {
	client := createDbConnection()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Panicf("Failed to close DB connection")
		}
	}()

	var result model.User

	dbName := os.Getenv("DB_NAME")
	dbUserCol := os.Getenv("DB_USER_COLLECTION")

	err := client.Database(dbName).Collection(dbUserCol).FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Panicf("Failed to retrieve document")
	}

	return result
}

func InsertOne(filter interface{}) {
	client := createDbConnection()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Panicf("Failed to close DB connection")
		}
	}()

	dbName := os.Getenv("DB_NAME")
	dbUserCol := os.Getenv("DB_USER_COLLECTION")

	_, err := client.Database(dbName).Collection(dbUserCol).InsertOne(context.TODO(), filter)

	if err != nil {
		log.Panicf("Failed to insert document")
	}
}
