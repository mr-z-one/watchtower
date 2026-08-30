package main

import (
	server "watchtower/Server"
	"watchtower/Server/Model/Job"
	"watchtower/Server/Model/db"
	"watchtower/rabbitmq"
)

func CreateIndexes() {
	Job.CreateIndex()
}

func main() {

	db.Mongo_connect()
	rabbitmq.Connect()
	CreateIndexes()

	//defer conn.Close()
	//defer ch.Close()

	server.StartServer(3000)
}
