package main

import (
	"lista/api/server"
	"lista/api/db"
	"log"
)

var version = "Version 1"
var name = "Lista-API"

func main() {
	m := db.NewMongoDb("db")
	m.Connect()
	log.Println("Starting", name, version)
	server.Start(":1337")
}
