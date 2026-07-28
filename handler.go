package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err == nil {
		jsonErr := Service.CreateUser(u)
		switch jsonErr {
		case InvalidData:
			http.Error(w, "JSON parsed but values are wrong", http.StatusUnprocessableEntity)
			return
		case UserExist:
			http.Error(w, "User with this ID already exist", http.StatusConflict)
			return
		default:
			rawJson, _ := json.Marshal(u)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write(rawJson)
			return
		}
	} else {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	if user, err := Service.GetUser(id["id"]); err == NotFound {
		http.Error(w, "Not Found", http.StatusNotFound)
	} else {
		rawJson, _ := json.Marshal(user)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(rawJson)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idURL := mux.Vars(r)
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err == nil {
		u.ID = idURL["id"]
		jsonErr := Service.UpdateUser(u)
		switch jsonErr {
		case NotFound:
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		case InvalidData:
			http.Error(w, "JSON parsed but values are wrong", http.StatusUnprocessableEntity)
			return
		default:
			rawJson, _ := json.Marshal(u)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(rawJson)
			return
		}
	} else {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idURL := mux.Vars(r)
	if err := Service.DeleteUser(idURL["id"]); err == NotFound {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
