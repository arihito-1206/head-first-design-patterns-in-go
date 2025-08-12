package main

type Beverage interface {
	Cost() float64
	GetDescription() string
}
