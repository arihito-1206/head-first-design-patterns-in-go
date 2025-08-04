package main

import "fmt"

type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
}

func (c *CurrentConditionsDisplay) Update(temp, humidity, pressure float64) {
	c.temperature = temp
	c.humidity = humidity
	c.Display()
}

func (c *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions: %.1f°C and %.1f%% humidity\n", c.temperature, c.humidity)
}
