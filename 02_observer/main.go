package main

func main() {
	weatherData := NewWeatherData()

	currentConditionsDisplay := &CurrentConditionsDisplay{}
	weatherData.RegsiterObsercer(currentConditionsDisplay)

	weatherData.SetMeasurements(25.5, 65.0, 1013.1)
	weatherData.SetMeasurements(27.0, 70.0, 1012.5)
}
