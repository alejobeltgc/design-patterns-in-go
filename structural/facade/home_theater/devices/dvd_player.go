package devices

import "fmt"

type DVDPlayer struct {
	Description string
	Movie       string
}

func NewDVDPlayer(description string) *DVDPlayer {
	return &DVDPlayer{
		Description: description,
		Movie:       "",
	}
}

func (d *DVDPlayer) On() {
	fmt.Printf("%s está encendido\n", d.Description)
}

func (d *DVDPlayer) Off() {
	fmt.Printf("%s está apagado\n", d.Description)
}

func (d *DVDPlayer) Play(movie string) {
	d.Movie = movie
	fmt.Printf("%s reproduciendo \"%s\"\n", d.Description, movie)
}

func (d *DVDPlayer) Stop() {
	fmt.Printf("%s detuvo \"%s\"\n", d.Description, d.Movie)
	d.Movie = ""
}

func (d *DVDPlayer) Pause() {
	fmt.Printf("%s pausó \"%s\"\n", d.Description, d.Movie)
}

func (d *DVDPlayer) Eject() {
	fmt.Printf("%s expulsó \"%s\"\n", d.Description, d.Movie)
	d.Movie = ""
}
