package listeners

import (
	"fmt"
)

type StatisticsDisplay struct {
	maxTemp     float64
	minTemp     float64
	tempSum     float64
	numReadings int
}

func NewStatisticsDisplay() *StatisticsDisplay {
	return &StatisticsDisplay{
		maxTemp: -999.0,
		minTemp: 999.0,
	}
}

func (sd *StatisticsDisplay) Update(temperature float64, humidity float64, pressure float64) {
	sd.tempSum += temperature
	sd.numReadings++

	if temperature > sd.maxTemp {
		sd.maxTemp = temperature
	}

	if temperature < sd.minTemp {
		sd.minTemp = temperature
	}

	sd.Display()
}

func (sd *StatisticsDisplay) Display() {
	avgTemp := sd.tempSum / float64(sd.numReadings)
	fmt.Printf("Avg/Max/Min temperature = %.1f/%.1f/%.1f\n",
		avgTemp, sd.maxTemp, sd.minTemp)
}
