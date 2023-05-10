package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"log"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]User{{
		Id: 1,
		Name: "Leonardo",
	}})
}

func main() {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/users", getUsers).Methods("GET")
	
	fmt.Println("api is on: 8080")
	log.Fatal(http.ListenAndServe(":8080", muxRouter))
}