package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nielsdingsbums/accessControl/structs"
	"net/http"
	"strconv"
)

var users = []structs.User{}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var newUser structs.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Printf("[ ERROR ] [createUser] %v\n", err)
		return
	}
	newUser.ID = len(users)
	fmt.Printf("[+] [createUser] %+v\n", newUser)
	users = append(users, newUser)

	json.NewEncoder(w).Encode(newUser)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if len(users)-1 >= id {
		fmt.Printf("[+] [getUser] id %d\n", id)
		json.NewEncoder(w).Encode(users[id])
	} else {
		fmt.Printf("[-] [getUser] id %d (index ouz of range)\n", id)
	}
}

