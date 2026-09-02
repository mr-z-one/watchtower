package program

import (
	handler "watchtower/Server/api/Program/Handler"

	"github.com/gorilla/mux"
)

func RegisterProgramApi(routes *mux.Router) {
	routes.HandleFunc("/program/all", handler.GetAllProgram)
	routes.HandleFunc("/program/name/{name}/", handler.GetProgramByName)
	routes.HandleFunc("/program/id/{id}/", handler.GetProgramByID)
}
