package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vojtech-mikes/opentrail/internals/db"
	"github.com/vojtech-mikes/opentrail/internals/model"
	"github.com/vojtech-mikes/opentrail/internals/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userEmail := r.URL.Query().Get("email")

	if userEmail == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}

	query := bson.D{{Key: "email", Value: userEmail}}

	user := db.FindOne(query)

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var userData model.User
	err := decoder.Decode(&userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}

	query := bson.D{
		{Key: "name", Value: userData.Name},
		{Key: "email", Value: userData.Email},
		{Key: "password", Value: utils.HashPassword(userData.Password)},
	}

	db.InsertOne(query)
}
