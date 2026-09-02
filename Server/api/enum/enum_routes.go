package enum

import (
	"watchtower/Server/api/enum/crtsh"
	subfiner "watchtower/Server/api/enum/subfinder"

	"github.com/gorilla/mux"
)

func RegisterEnumerationRoutes(routes *mux.Router) {
	enum := routes.PathPrefix("/enum").Subrouter()
	crtsh.RegisterCrtshApi(enum)
	subfiner.RegisterSubfinderApi(enum)
}
