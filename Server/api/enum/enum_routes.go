package program

import (
	"watchtower/Server/api/enum/crtsh"

	"github.com/gorilla/mux"
)

func RegisterEnumerationRoutes(routes *mux.Router) {
	enum := routes.PathPrefix("/enum").Subrouter()
	crtsh.RegisterCrtshApi(enum)
}
