package handler

import (
	"net/http"
	"watchtower/Server/Model/program"
	"watchtower/config"
)

func GetAllSyncProgram(w http.ResponseWriter, r *http.Request) {
	programConfig := config.ReadeConfig("config.json")

	program.UpsertProgram(programConfig)
	w.Write([]byte("sync is successful!"))
}
