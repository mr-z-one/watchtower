package main

import (
	server "watchtower/Server"
	"watchtower/Server/Model/db"
)

func main() {

	db.Mongo_connect()
	server.StartServer(3000)
}
