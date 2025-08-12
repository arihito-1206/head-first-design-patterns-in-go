package main

type Espresso struct {
	description string
}

func (e *Espresso) Cost() float64 {
	return 1.99
}
func (e *Espresso) GetDescription() string {
	return e.description
}

func NewEspresso() *Espresso {
	return &Espresso{description: "Espresso"}
}
