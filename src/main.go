package main

import (
	"lista/api/server"
	"log"
)

var version = "Version 1"
var name = "Lista-API"

func main() {
	log.Println("Starting", name, version)
	server.Start(":1337")
}
