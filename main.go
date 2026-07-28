package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"restAPI/handlers"
	"restAPI/service"
	"restAPI/storage"

	"github.com/gorilla/mux"
)

func main() {
	stor := storage.NewMapStorage()
	service := service.NewUserService(stor)
	handler := handlers.NewHandler(service)

	connStr := "user=postgres password=21282908 dbname=usersdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	fmt.Println(db, err)

	mux := mux.NewRouter()
	mux.HandleFunc("/users", handler.CreateUser).Methods("POST")
	mux.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")
	mux.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")
	mux.HandleFunc("/users/{id}", handler.DeleteUser).Methods("DELETE")

	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)
	http.ListenAndServe(port, mux)
}
