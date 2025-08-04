package main

type WeatherData struct {
	observers   []Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (w *WeatherData) RegsiterObsercer(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o Observer) {
	for i, obs := range w.observers {
		if obs == o {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *WeatherData) NotifiyObservers() {
	for _, o := range w.observers {
		o.Update(w.temperature, w.humidity, w.pressure)
	}
}

func (w *WeatherData) SetMeasurements(temp, humidity, pressure float64) {
	w.temperature = temp
	w.humidity = humidity
	w.pressure = pressure
	w.NotifiyObservers()
}
