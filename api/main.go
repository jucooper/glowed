package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// RocketLeagueStatsPlayerEndpoint is for retrieving a single player
const RocketLeagueStatsPlayerEndpoint = "https://api.rocketleaguestats.com/v1/player"

// Player Struct from API
type Player struct {
	DisplayName string `json:"displayName"`
	Stats       `json:"stats"`
}

// Stats struct from API
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

	client := &http.Client{}

	// Init the request
	req, err := http.NewRequest("GET", RocketLeagueStatsPlayerEndpoint, nil)
	if err != nil {
		log.Print(err)
	}

	// Append params to request
	q := req.URL.Query()
	q.Add("apikey", os.Getenv("RLS_API_KEY"))
	q.Add("unique_id", os.Getenv("PROFILE"))
	q.Add("platform_id", os.Getenv("PLATFORM"))

	req.URL.RawQuery = q.Encode()

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
	}

	// Decode the response
	var data Player
	err = DecodeResp(resp, &data)
	if err != nil {
		log.Print(err)
	}

	json, _ := json.Marshal(Response{Player{data.DisplayName, Stats{data.Wins, data.Goals, data.Mvps, data.Saves, data.Shots, data.Assists}}})
	w.Write(json)
}

// DecodeResp decodes the MapBox response and stores in the related struct
func DecodeResp(resp *http.Response, data *Player) error {
	err := json.NewDecoder(resp.Body).Decode(data)
	return err
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootRequest).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}
