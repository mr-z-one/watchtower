package program

import (
	"fmt"

	"github.com/gorilla/mux"
)

func RegisterEnumerationRoutes(routes *mux.Router) {
	enum := routes.PathPrefix("/enum").Subrouter()
	fmt.Println(enum)
}
