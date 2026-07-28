package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var Service *UserService = NewUserService()

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/users", CreateUser).Methods("POST")
	mux.HandleFunc("/users/{id}", GetUser).Methods("GET")
	mux.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	mux.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)
	http.ListenAndServe(port, mux)
}
