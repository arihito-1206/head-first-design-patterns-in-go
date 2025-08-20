package main

import "fmt"

type NYPizzaStore struct{}

func (s NYPizzaStore) CreatePizza(pizzaType string) Pizza {
	if pizzaType == "cheese" {
		return NewNYStyleCheesePizza()
	}
	return nil
}

type NYStyleCheesePizza struct {
	name string
}

func NewNYStyleCheesePizza() *NYStyleCheesePizza {
	return &NYStyleCheesePizza{name: "NY Style Cheese Pizza"}
}

func (p *NYStyleCheesePizza) GetName() string { return p.name }
func (p *NYStyleCheesePizza) Prepare()        { fmt.Println("Prepare", p.name) }
func (p *NYStyleCheesePizza) Bake()           { fmt.Println("Bake", p.name) }
func (p *NYStyleCheesePizza) Cut()            { fmt.Println("Cut", p.name) }
func (p *NYStyleCheesePizza) Box()            { fmt.Println("Box", p.name) }
