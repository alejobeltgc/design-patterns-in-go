package publisher

import (
	"designpatterns/behavioral/observer/weather/listeners"
)

type WeatherData struct {
	observerList []listeners.WeatherListener
	temperature  float64
	humidity     float64
	pressure     float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (wd *WeatherData) RegisterObserver(o listeners.WeatherListener) {
	wd.observerList = append(wd.observerList, o)
}

func (wd *WeatherData) RemoveObserver(o listeners.WeatherListener) {
	for i, observer := range wd.observerList {
		if observer == o {
			wd.observerList = append(wd.observerList[:i], wd.observerList[i+1:]...)
			break
		}
	}

}

func (wd *WeatherData) NotifyObservers() {
	for _, observer := range wd.observerList {
		observer.Update(wd.temperature, wd.humidity, wd.pressure)
	}
}

func (wd *WeatherData) SetMeasurements(tmp float64, hum float64, pre float64) {
	wd.temperature = tmp
	wd.humidity = hum
	wd.pressure = pre
	wd.MeasurementsChanged()
}

func (wd *WeatherData) GetTemperature() float64 {
	return wd.temperature
}

func (wd *WeatherData) GetHumidity() float64 {
	return wd.humidity
}

func (wd *WeatherData) GetPressure() float64 {
	return wd.pressure
}

func (wd *WeatherData) MeasurementsChanged() {
	wd.NotifyObservers()
}
