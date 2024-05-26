package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"qayyuum/go_fintech/users"

	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {

	// Check if the body is in order
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to parse request", http.StatusBadRequest)
		return
	}
	// Check for login info
	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	if err != nil {
		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		return
	}

	login, err := users.Login(formattedBody.Username, formattedBody.Password)
	if err != nil {
		http.Error(w, "Unable to login", http.StatusBadRequest)
		return
	}

	// Create response
	if login["messages"] == "Login successful" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

// StartAPI - Create an API
func StartAPI() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	fmt.Println("App is working in port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
