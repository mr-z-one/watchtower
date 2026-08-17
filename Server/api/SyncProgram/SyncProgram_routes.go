package syncprogram

import (
	handler "watchtower/Server/api/SyncProgram/Handler"

	"github.com/gorilla/mux"
)

func RegisterSyncProgramApi(routes *mux.Router) {
	routes.HandleFunc("/sync", handler.GetAllSyncProgram)
}
