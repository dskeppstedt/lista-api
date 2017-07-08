package main

import (
	"lista/api/db"
	"lista/api/server"
	"log"
)

var version = "Version 1"
var name = "Lista-API"

//Global accesse to database
var DbM *db.Mongodb

func main() {

	DbM := db.NewMongoDb("db")
	DbM.Connect()
	server.DbStore = DbM

	log.Println("Starting", name, version)
	server.Start(":1337")
}
