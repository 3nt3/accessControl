package api

import "net/http"

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("lol"))
}
