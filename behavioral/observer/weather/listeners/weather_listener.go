package listeners

type WeatherListener interface {
	Update(temperature float64, humidity float64, pressure float64)
}
