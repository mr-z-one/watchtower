package subdomain

import (
	handler "watchtower/Server/api/subdomain/Handler"

	"github.com/gorilla/mux"
)

func RegisterSubdomainApi(routes *mux.Router) {
	routes.HandleFunc("/subdomain/add", handler.AddSubdomain).Methods("POST")

}
