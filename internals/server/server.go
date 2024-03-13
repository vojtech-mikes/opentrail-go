package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vojtech-mikes/opentrail/internals/handlers"
)

func initHandlers(r *mux.Router) {
	r.HandleFunc("/user", handlers.GetUser).Methods("GET")
	r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
}

func InitServer(port string) {
	r := mux.NewRouter()
	initHandlers(r)
	http.ListenAndServe(port, r)
}
