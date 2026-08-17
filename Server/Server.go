package server

import (
	"log"
	"net/http"
	"strconv"
	"watchtower/Server/Model/db"
	"watchtower/Server/api"
	custom_color "watchtower/custom_Color"

	"github.com/gorilla/mux"
)

func StartServer(port int) {

	db.Mongo_connect()

	r := mux.NewRouter().StrictSlash(true)
	api.RegisterApi(r)

	custom_color.Succeed()("[+] Server start at http://localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))

}
