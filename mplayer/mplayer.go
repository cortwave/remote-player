package mplayer

import (
	"io"
	"log"
	"os/exec"
)

var stdin io.WriteCloser
var cmd *exec.Cmd

//PlayerMessage for playing manipulations
type PlayerMessage interface {
	Name() string
}

//PauseMessage for play/pause
type PauseMessage struct {
}

//Name of PauseMessage
func (p PauseMessage) Name() string {
	return "pause"
}

//NewSongMessage for new song playing
type NewSongMessage struct {
	URL string
}

//Name of NewSongMessage
func (n NewSongMessage) Name() string {
	return "newSong"
}

//QuitMessage for quit
type QuitMessage struct {
}

//Name of QuitMessage
func (q QuitMessage) Name() string {
	return "quit"
}

//IncreaseVolumeMessage for volume increasing
type IncreaseVolumeMessage struct {
	Points int
}

//Name of IncreaseVolumeMessage
func (i IncreaseVolumeMessage) Name() string {
	return "increaseVolume"
}

//DecreaseVolumeMessage for volume decreasing
type DecreaseVolumeMessage struct {
	Points int
}

//Name of DecreaseVolumeMessage
func (i DecreaseVolumeMessage) Name() string {
	return "decreaseVolume"
}

//StartPlayer returns channel for commands
func StartPlayer() chan<- PlayerMessage {
	messagesChannel := make(chan PlayerMessage)
	go handleMessages(messagesChannel)
	return messagesChannel
}

func handleMessages(messagesChannel <-chan PlayerMessage) {
	for {
		message := <-messagesChannel
		switch message.(type) {
		case PauseMessage:
			pause()
		case NewSongMessage:
			log.Println(message)
			go playNew(message.(NewSongMessage).URL)
		case QuitMessage:
			quitPlayer()
		case IncreaseVolumeMessage:
			increaseVolume(message.(IncreaseVolumeMessage).Points)
		case DecreaseVolumeMessage:
			decreaseVolume(message.(DecreaseVolumeMessage).Points)
		}
	}
}

func pause() {
	writeToPlayer("p")
}

func playNew(url string) {
	if cmd != nil {
		quitPlayer()
	}
	cmd = playCmd(url)
	stdinPlayer, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	stdin = stdinPlayer
	executeCmd(cmd)
}

func quitPlayer() {
	writeToPlayer("q")
}

func increaseVolume(points int) {
	for i := 0; i < points; i++ {
		writeToPlayer("*")
	}
}

func decreaseVolume(points int) {
	for i := 0; i < points; i++ {
		writeToPlayer("/")
	}
}

func writeToPlayer(message string) {
	if stdin != nil {
		io.WriteString(stdin, message)
	}
}

func playCmd(url string) *exec.Cmd {
	return exec.Command("mplayer", url)
}

func executeCmd(cmd *exec.Cmd) {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
