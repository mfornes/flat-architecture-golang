package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		users, err := db.FindUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(users)
		return
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ID, err := strconv.Atoi(r.URL.Path[len("/users/"):])
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid User id", r.URL.Path[len("/users/"):]), http.StatusBadRequest)
			return
		}

		result, err := db.FindUser(ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
		return
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var newUser User
		err := decoder.Decode(&newUser)
		if err != nil {
			http.Error(w, "error while parsing new User data: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := db.SaveUser(newUser); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("success")
		return
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		ID, err := strconv.Atoi(r.URL.Path[len("/users/delete/"):])
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid User id", r.URL.Path[len("/users/delete/"):]), http.StatusBadRequest)
			return
		}

		err = db.DeleteUser(ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("success")
		return
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}
