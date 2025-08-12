package main

type Mocha struct {
	beverage Beverage
}

func NewMocha(bev Beverage) *Mocha {
	return &Mocha{bev}
}

func (m *Mocha) GetDescription() string {
	return m.beverage.GetDescription() + " + Mocha"
}

func (m *Mocha) Cost() float64 {
	return m.beverage.Cost() + 0.20
}
