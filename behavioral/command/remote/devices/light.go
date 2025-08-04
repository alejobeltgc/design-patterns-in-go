package devices

import "fmt"

type Light struct {
	location string
}

func NewLight(location string) *Light {
	return &Light{location: location}
}

func (l *Light) On() {
	fmt.Printf("Luz de %s encendida\n", l.location)
}

func (l *Light) Off() {
	fmt.Printf("Luz de %s apagada\n", l.location)
}
