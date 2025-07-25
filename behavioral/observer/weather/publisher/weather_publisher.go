package publisher

type WeatherPublisher interface {
	RegisterObserver()
	RemoveObserver()
	NotifyObservers()
}
