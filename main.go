package main

import (
	"github.com/gorilla/mux"
	"github.com/nielsdingsbums/accessControl/api"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/getUsers/", api.GetUsers)

	http.ListenAndServe(":8080", nil)
}
