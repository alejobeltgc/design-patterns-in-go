package publisher

import "designpatterns/behavioral/observer/weather/listeners"

type WeatherPublisher interface {
	RegisterObserver(o listeners.WeatherListener)
	RemoveObserver(o listeners.WeatherListener)
	NotifyObservers()
}
