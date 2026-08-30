package main

import (
	server "watchtower/Server"
	"watchtower/Server/Model/db"
)

func main() {

	db.Mongo_connect()
	db.InitializeIndexes()
	//defer conn.Close()
	//defer ch.Close()

	server.StartServer(3000)
}
