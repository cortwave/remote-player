package main

import (
	"encoding/json"
	"log"
	"net/http"
	"remote-player/mplayer"
)

type songRequest struct {
	URL string `json:"url"`
}

type changeVolumeRequest struct {
	Points int `json:"points"`
}

var playerMessages = mplayer.StartPlayer()

func main() {
	startHTTPServer()
}

func startHTTPServer() {
	http.HandleFunc("/play", handlePlay)
	http.HandleFunc("/quit", handleQuit)
	http.HandleFunc("/increase", handleIncreaseVolume)
	http.HandleFunc("/decrease", handleDecreaseVolume)
	http.HandleFunc("/add", handleAddSongMessage)
	http.ListenAndServe(":8000", nil)
}

func handlePlay(w http.ResponseWriter, r *http.Request) {
	log.Println("play")
	playerMessages <- mplayer.PauseMessage{}
}

func handleQuit(w http.ResponseWriter, r *http.Request) {
	log.Println("quit")
	playerMessages <- mplayer.QuitMessage{}
}

func handleIncreaseVolume(w http.ResponseWriter, r *http.Request) {
	log.Println("increase volume")
	request := changeVolumeRequest{}
	json.NewDecoder(r.Body).Decode(&request)
	increaseVolumeMessage := mplayer.IncreaseVolumeMessage{Points: request.Points}
	playerMessages <- increaseVolumeMessage
}

func handleDecreaseVolume(w http.ResponseWriter, r *http.Request) {
	log.Println("decrease volume")
	request := changeVolumeRequest{}
	json.NewDecoder(r.Body).Decode(&request)
	decreaseVolumeMessage := mplayer.DecreaseVolumeMessage{Points: request.Points}
	playerMessages <- decreaseVolumeMessage
}

func handleAddSongMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("add song")
	request := songRequest{}
	json.NewDecoder(r.Body).Decode(&request)
	addSongMessage := mplayer.AddSongMessage{URL: request.URL}
	playerMessages <- addSongMessage
}
