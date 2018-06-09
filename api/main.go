package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Player Struct
type Player struct {
	DisplayName string `json:"displayName"`
	Stats       `json:"stats"`
}

// Stats Struct
type Stats struct {
	Wins    int `json:"wins"`
	Goals   int `json:"goals"`
	Mvps    int `json:"mvps"`
	Saves   int `json:"saves"`
	Shots   int `json:"shots"`
	Assists int `json:"assists"`
}

// Response struct for service response
type Response struct {
	Player `json:"player"`
}

// RootRequest handles the root endpoint
func RootRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json, _ := json.Marshal(Response{Player{"Test", Stats{1, 2, 3, 4, 5, 6}}})
	w.Write(json)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootRequest).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}
