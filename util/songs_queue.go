package util

import (
	"container/list"
	"log"
)

//Song with info
type Song struct {
	URL         string
	Description string
}

var songs = list.New()
var currentSong *list.Element
var lock = make(semaphore, 999)

//Push to queue
func Push(song Song) {
	songs.PushBack(song)
	lock.Unlock()
	log.Println("Push end")
}

//Next song
func Next() Song {
	var nextSong *list.Element
	if currentSong == nil {
		nextSong = songs.Front()
	} else {
		nextSong = currentSong.Next()
	}
	if nextSong == nil {
		lock.Lock()
		return Next()
	}
	currentSong = nextSong
	return currentSong.Value.(Song)
}

//IsEmpty checks if songs queue is empty
func IsEmpty() bool {
	return songs.Front() == nil
}
