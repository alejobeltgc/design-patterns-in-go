package main

import (
	"designpatterns/behavioral/observer/weather/listeners"
	"designpatterns/behavioral/observer/weather/publisher"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Weather Station - Observer Pattern Demo ===")

	weatherData := publisher.NewWeatherData()

	currentDisplay := listeners.NewCurrentConditionsDisplay()
	statisticsDisplay := listeners.NewStatisticsDisplay()
	forecastDisplay := listeners.NewForecastDisplay()

	// Registrar los observers
	weatherData.RegisterObserver(currentDisplay)
	weatherData.RegisterObserver(statisticsDisplay)
	weatherData.RegisterObserver(forecastDisplay)

	fmt.Println("\n1. Primera medición:")
	fmt.Println(strings.Repeat("-", 40))
	weatherData.SetMeasurements(26.6, 65.0, 1013.1)

	fmt.Println("\n2. Segunda medición:")
	fmt.Println(strings.Repeat("-", 40))
	weatherData.SetMeasurements(27.7, 70.0, 997.0)

	fmt.Println("\n3. Tercera medición:")
	fmt.Println(strings.Repeat("-", 40))
	weatherData.SetMeasurements(25.5, 90.0, 1005.0)

	fmt.Println("\n4. Removiendo el forecast display:")
	fmt.Println(strings.Repeat("-", 40))
	weatherData.RemoveObserver(forecastDisplay)
	weatherData.SetMeasurements(24.2, 85.0, 1008.5)

	fmt.Println("\n=== Demo completado ===")
}
