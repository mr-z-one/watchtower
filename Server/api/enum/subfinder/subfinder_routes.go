package subfiner

import (
	handler "watchtower/Server/api/enum/subfinder/Handler"

	"github.com/gorilla/mux"
)

func RegisterSubfinderApi(routes *mux.Router) {

	routes.HandleFunc("/subfinder", handler.SendSubFinderJob).Methods("POST")

}
