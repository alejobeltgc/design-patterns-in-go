package listeners

import (
	"fmt"
)

type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
	pressure    float64
}

func NewCurrentConditionsDisplay() *CurrentConditionsDisplay {
	return &CurrentConditionsDisplay{}
}

func (ccd *CurrentConditionsDisplay) Update(temperature float64, humidity float64, pressure float64) {
	ccd.temperature = temperature
	ccd.humidity = humidity
	ccd.pressure = pressure
	ccd.Display()
}

func (ccd *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions: %.1f°C and %.1f%% humidity, %.1f pressure\n",
		ccd.temperature, ccd.humidity, ccd.pressure)
}
