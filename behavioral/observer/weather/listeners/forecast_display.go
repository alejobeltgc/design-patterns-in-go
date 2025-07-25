package listeners

import (
	"fmt"
)

type ForecastDisplay struct {
	currentPressure float64
	lastPressure    float64
}

func NewForecastDisplay() *ForecastDisplay {
	return &ForecastDisplay{}
}

func (fd *ForecastDisplay) Update(temperature float64, humidity float64, pressure float64) {
	fd.lastPressure = fd.currentPressure
	fd.currentPressure = pressure
	fd.Display()
}

func (fd *ForecastDisplay) Display() {
	fmt.Print("Forecast: ")
	if fd.currentPressure > fd.lastPressure {
		fmt.Println("Improving weather on the way!")
	} else if fd.currentPressure == fd.lastPressure {
		fmt.Println("More of the same")
	} else if fd.currentPressure < fd.lastPressure {
		fmt.Println("Watch out for cooler, rainy weather")
	}
}
