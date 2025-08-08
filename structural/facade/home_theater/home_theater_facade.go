package main

import (
	"designpatterns/structural/facade/home_theater/devices"
	"fmt"
)

type HomeTheaterFacade struct {
	amp       *devices.Amplifier
	tuner     *devices.Tuner
	dvd       *devices.DVDPlayer
	projector *devices.Projector
	screen    *devices.Screen
	lights    *devices.TheaterLights
	popper    *devices.PopcornPopper
}

func NewHomeTheaterFacade(
	amp *devices.Amplifier,
	tuner *devices.Tuner,
	dvd *devices.DVDPlayer,
	projector *devices.Projector,
	screen *devices.Screen,
	lights *devices.TheaterLights,
	popper *devices.PopcornPopper,
) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		amp:       amp,
		tuner:     tuner,
		dvd:       dvd,
		projector: projector,
		screen:    screen,
		lights:    lights,
		popper:    popper,
	}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Preparándose para ver película...")

	h.popper.On()
	h.popper.Pop()

	h.lights.Dim(10)

	h.screen.Down()

	h.projector.On()
	h.projector.WideScreenMode()

	h.amp.On()
	h.amp.SetSurroundSound()
	h.amp.SetVolume(5)

	h.dvd.On()
	h.dvd.Play(movie)

	fmt.Printf("¡Disfruta de la película \"%s\"!\n", movie)
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Apagando el home theater...")

	h.popper.Off()

	h.lights.On()

	h.screen.Up()

	h.projector.Off()

	h.amp.Off()

	h.dvd.Stop()
	h.dvd.Eject()
	h.dvd.Off()

	fmt.Println("¡Home theater apagado!")
}

func (h *HomeTheaterFacade) ListenToRadio(frequency float64) {
	fmt.Printf("Sintonizando radio a %.1f FM...\n", frequency)

	h.tuner.On()
	h.tuner.SetFM()
	h.tuner.SetFrequency(frequency)

	h.amp.On()
	h.amp.SetStereoSound()
	h.amp.SetVolume(3)

	fmt.Printf("¡Radio sintonizada a %.1f FM!\n", frequency)
}

func (h *HomeTheaterFacade) EndRadio() {
	fmt.Println("Apagando radio...")

	h.tuner.Off()

	h.amp.Off()

	fmt.Println("¡Radio apagada!")
}
