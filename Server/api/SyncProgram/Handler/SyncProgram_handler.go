package handler

import (
	"fmt"
	"net/http"
	"watchtower/Server/Model/Job"
	"watchtower/Server/Model/program"
	"watchtower/config"
	"watchtower/rabbitmq"
)

func GetAllSyncProgram(w http.ResponseWriter, r *http.Request) {
	programConfig := config.ReadeConfig("config.json")
	m := rabbitmq.NewMessage("crtssh", "sess.sku.ac.ir")
	go rabbitmq.SendMessage(m, func(ack bool) { fmt.Println("acked From Synced: ", ack) })
	Job.Insert(Job.PENDING, m)

	program.UpsertProgram(programConfig)
	w.Write([]byte("sync is successful!"))
}
