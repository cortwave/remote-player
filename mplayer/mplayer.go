package mplayer

import (
	"io"
	"log"
	"os/exec"
)

var stdin io.WriteCloser
var cmd *exec.Cmd

//StartPlayer returns channel for commands
func StartPlayer() chan<- string {
	messagesChannel := make(chan string)
	go handleMessages(messagesChannel)
	return messagesChannel
}

func handleMessages(messagesChannel <-chan string) {
	for {
		message := <-messagesChannel
		switch message {
		case "p":
			writeToPlayer(message)
		default:
			log.Println(message)
			go playNew(message)
		}
	}
}

func writeToPlayer(message string) {
	if stdin != nil {
		io.WriteString(stdin, message)
	}
}

func playNew(url string) {
	if cmd != nil {
		writeToPlayer("q")
	}
	cmd = playCmd(url)
	stdinPlayer, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	stdin = stdinPlayer
	go executeCmd(cmd)
}

func playCmd(url string) *exec.Cmd {
	return exec.Command("mplayer", url)
}

func executeCmd(cmd *exec.Cmd) {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
