package main

type HouseBlend struct {
	description string
}

func (h *HouseBlend) GetDescription() string {
	return h.description
}

func (h *HouseBlend) Cost() float64 {
	return 0.89
}

func NewHouseBlend() *HouseBlend {
	return &HouseBlend{description: "HouseBlend Coffee"}
}
