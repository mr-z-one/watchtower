package handler

import (
	"encoding/json"
	"net/http"
	"watchtower/Server/Message"
	"watchtower/Server/Model/program"

	"github.com/gorilla/mux"
)

func GetAllProgram(w http.ResponseWriter, r *http.Request) {

	all_program, err := program.GetAllProgram()
	if err != nil {
		w.Write([]byte("some error.."))
	}
	data, _ := json.Marshal(all_program)
	w.Write(data)
}

func GetProgramByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	current_program, err := program.GetProgramByName(name)
	if err != nil {
		data, _ := json.Marshal(Message.CreateErrorMessage("no result", 404))
		w.Write(data)
		return
	}
	data, _ := json.Marshal(current_program)
	w.Write(data)
}
