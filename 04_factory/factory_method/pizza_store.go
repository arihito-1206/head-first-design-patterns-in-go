package main

import "fmt"

type Pizza interface {
	GetName() string
	Prepare()
	Bake()
	Cut()
	Box()
}

type PizzaStore interface {
	CreatePizza(pizzatype string) Pizza
}

func OrderPizza(store PizzaStore, pizzaType string) Pizza {
	pizza := store.CreatePizza(pizzaType)
	fmt.Println("--- Making a", pizza.GetName(), "---")

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}
