package devices

import "fmt"

type TheaterLights struct {
	Description string
	Brightness  int
}

func NewTheaterLights(description string) *TheaterLights {
	return &TheaterLights{
		Description: description,
		Brightness:  100,
	}
}

func (t *TheaterLights) On() {
	t.Brightness = 100
	fmt.Printf("%s están encendidas\n", t.Description)
}

func (t *TheaterLights) Off() {
	t.Brightness = 0
	fmt.Printf("%s están apagadas\n", t.Description)
}

func (t *TheaterLights) Dim(level int) {
	t.Brightness = level
	fmt.Printf("%s atenuadas al %d%%\n", t.Description, level)
}
