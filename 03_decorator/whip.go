package main

type Whip struct {
	beverage Beverage
}

func NewWhip(bev Beverage) *Whip {
	return &Whip{bev}
}

func (w *Whip) GetDescription() string {
	return w.beverage.GetDescription() + " + Whip"
}

func (w *Whip) Cost() float64 {
	return w.beverage.Cost() + 0.10
}
