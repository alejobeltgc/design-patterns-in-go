package concretedecorator

import "designpatterns/structural/decorator/starbuzz/component"

type Milk struct {
	Beverage component.Beverage
}

func NewMilk(b component.Beverage) *Milk {
	if b == nil {
		panic("beverage cannot be nil")
	}
	return &Milk{
		Beverage: b,
	}
}

func (m *Milk) Cost() float64 {
	return m.Beverage.Cost() + .10
}

func (m *Milk) GetDescription() string {
	return m.Beverage.GetDescription() + ", Milk"
}
