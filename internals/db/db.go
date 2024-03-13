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

	envs := utils.PrepareDbEnvs()

	var result model.User

	err := client.Database(envs.DbName).Collection(envs.DbUserCol).FindOne(context.TODO(), filter).Decode(&result)

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

	envs := utils.PrepareDbEnvs()

	_, err := client.Database(envs.DbName).Collection(envs.DbUserCol).InsertOne(context.TODO(), filter)

	if err != nil {
		log.Panicf("Failed to insert document")
	}
}
