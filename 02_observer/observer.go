package main

type Observer interface {
	Update(tenp float64, humidity float64, pressure float64)
}
