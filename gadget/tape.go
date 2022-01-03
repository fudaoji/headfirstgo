package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t *TapePlayer) Play(song string) {
	fmt.Println("palying", song)
}

func (t *TapePlayer) Stop() {
	fmt.Println("Stoped!")
}

type TapeRecoder struct {
	Microphones int
}

func (t *TapeRecoder) Play(song string) {
	fmt.Println("Palying", song)
}

func (t *TapeRecoder) Record() {
	fmt.Println("Recording!")
}

func (t *TapeRecoder) Stop() {
	fmt.Println("Stoped!")
}

func playList(device *TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TestGadget() {
	songs := []string{"Yesterday once more", "Take me to your heart"}
	player := TapePlayer{}
	playList(&player, songs)
}
