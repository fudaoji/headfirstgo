package main

import (
	"github.com/fudaoji/headfirstgo/gadget"
)

func playList(device *gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	songs := []string{"Yesterday once more", "Take me to your heart"}
	player := gadget.TapePlayer{}
	playList(player, songs)
}
