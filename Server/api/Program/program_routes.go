package program

import (
	handler "watchtower/Server/api/Program/Handler"

	"github.com/gorilla/mux"
)

func RegisterProgramApi(routes *mux.Router) {
	routes.HandleFunc("/program/all", handler.GetAllProgram)
	routes.HandleFunc("/program/{name}/", handler.GetProgramByName)
}
