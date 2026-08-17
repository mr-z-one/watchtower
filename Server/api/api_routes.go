package api

import (
	program "watchtower/Server/api/Program"
	syncProgram "watchtower/Server/api/SyncProgram"
	enum "watchtower/Server/api/enum"

	"github.com/gorilla/mux"
)

func RegisterApi(routes *mux.Router) {
	api_routes := routes.PathPrefix("/api").Subrouter()
	program.RegisterProgramApi(api_routes)

	syncProgram.RegisterSyncProgramApi(api_routes)
	enum.RegisterEnumerationRoutes(api_routes)
}
