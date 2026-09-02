package handler

import (
	"encoding/json"
	"net/http"
	"watchtower/Server/Message"
	"watchtower/Server/Model/program"
	"watchtower/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllProgram(w http.ResponseWriter, r *http.Request) {

	all_program, err := program.GetAllProgram()
	e := utils.FailOnError(err, "", nil)
	if e {

		w.Write([]byte("some error.."))
	}

	data, _ := json.Marshal(all_program)
	w.Write(data)
}

func GetProgramByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objectId, err := primitive.ObjectIDFromHex(id)
	if utils.FailOnError(err, "", nil) {

		data, _ := json.Marshal(Message.CreateErrorMessage("id is not valid", 400))
		w.Write(data)
		return
	}
	current_program, err := program.GetProgramByID(objectId)

	if utils.FailOnError(err, "", nil) {

		data, _ := json.Marshal(Message.CreateErrorMessage("no result", 404))
		w.Write(data)
		return
	}

	data, _ := json.Marshal(current_program)
	w.Write(data)
}

func GetProgramByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	current_program, err := program.GetProgramByName(name)

	e := utils.FailOnError(err, "", nil)
	if e {

		data, _ := json.Marshal(Message.CreateErrorMessage("no result", 404))
		w.Write(data)
		return
	}

	data, _ := json.Marshal(current_program)
	w.Write(data)
}
