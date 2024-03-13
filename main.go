package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vojtech-mikes/opentrail/internals/server"
)

func main() {

	err := godotenv.Load("local.env")

	if err != nil {
		log.Panicf("ERROR loading env variables")
	}

	port := os.Getenv("PORT")

	server.InitServer(port)
}
