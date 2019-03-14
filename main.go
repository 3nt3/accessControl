package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nielsdingsbums/accessControl/api"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/getUsers", api.GetUsers)
	r.HandleFunc("/createUser", api.CreateUser)
	r.HandleFunc("/getUser/{id}", api.GetUser)
	r.HandleFunc("/hasAccess/{id}", api.HasAccess)

	fmt.Print(http.ListenAndServe(":8000", r))
}
