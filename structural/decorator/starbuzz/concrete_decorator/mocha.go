package concretedecorator

import "designpatterns/structural/decorator/starbuzz/component"

type Mocha struct {
	Beverage component.Beverage
}

func NewMocha(b component.Beverage) *Mocha {
	if b == nil {
		panic("beverage cannot be nil")
	}
	return &Mocha{
		Beverage: b,
	}
}

func (m *Mocha) Cost() float64 {
	return m.Beverage.Cost() + .20
}

func (m *Mocha) GetDescription() string {
	return m.Beverage.GetDescription() + ", Mocha"
}
