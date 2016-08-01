package main

import (
	"encoding/json"
	"log"
	"net/http"
	"remote-player/mplayer"
)

type newRequest struct {
	URL string `json:"url"`
}

var playerMessages = mplayer.StartPlayer()

func main() {
	startHTTPServer()
}

func startHTTPServer() {
	http.HandleFunc("/play", handlePlay)
	http.HandleFunc("/new", handleNew)
	http.ListenAndServe(":8000", nil)
}

func handlePlay(w http.ResponseWriter, r *http.Request) {
	log.Println("play")
	playerMessages <- "p"
}

func handleNew(w http.ResponseWriter, r *http.Request) {
	log.Println("new")
	request := newRequest{}
	json.NewDecoder(r.Body).Decode(&request)
	playerMessages <- request.URL
}
