package utils

import (
	"log"
	"os"

	"github.com/vojtech-mikes/opentrail/internals/model"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pass), 8)

	return string(hashed)
}

func PrepareDbEnvs() model.DbEnvs {

	dbName := os.Getenv("DB_NAME")
	dbUserCol := os.Getenv("DB_USER_COLLECTION")

	if dbName == "" {
		log.Panicf("Failed to get DB_NAME")
	}

	if dbUserCol == "" {
		log.Panicf("Failed to get DB_USER_COLLECTION")
	}

	return model.DbEnvs{DbName: dbName, DbUserCol: dbUserCol}
}
