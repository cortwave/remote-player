package mplayer

import (
	"io"
	"log"
	"os/exec"
	"remote-player/util"
)

var stdin io.WriteCloser
var cmd *exec.Cmd

//StartPlayer returns channel for commands
func StartPlayer() chan<- PlayerMessage {
	messagesChannel := make(chan PlayerMessage)
	go handleMessages(messagesChannel)
	go startPlaying()
	return messagesChannel
}

func handleMessages(messagesChannel <-chan PlayerMessage) {
	for {
		message := <-messagesChannel
		switch message.(type) {
		case PauseMessage:
			pause()
		case QuitMessage:
			quitPlayer()
		case IncreaseVolumeMessage:
			increaseVolume(message.(IncreaseVolumeMessage).Points)
		case DecreaseVolumeMessage:
			decreaseVolume(message.(DecreaseVolumeMessage).Points)
		case AddSongMessage:
			song := util.Song{URL: message.(AddSongMessage).URL}
			util.Push(song)
		}
	}
}

func pause() {
	writeToPlayer("p")
}

func startPlaying() {
	url := util.Next().URL
	log.Println(url)
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
	log.Println("play end")
	startPlaying()
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
