package mplayer

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

//AddSongMessage to playlist
type AddSongMessage struct {
	URL string
}

//Name of AddSongMessage
func (a AddSongMessage) Name() string {
	return "addSong"
}
