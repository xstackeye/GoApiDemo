package main

import (
	"GoApiDemo/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.Methods("GET").Path("/users").HandlerFunc(handlers.UserHandler)
	r.Methods("POST").Path("/users").HandlerFunc(handlers.AddUserHandler)
	return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}
