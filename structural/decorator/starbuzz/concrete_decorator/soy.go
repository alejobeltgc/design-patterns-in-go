package concretedecorator

import "designpatterns/structural/decorator/starbuzz/component"

type Soy struct {
	Beverage component.Beverage
}

func NewSoy(b component.Beverage) *Soy {
	if b == nil {
		panic("beverage cannot be nil")
	}
	return &Soy{
		Beverage: b,
	}
}

func (s *Soy) Cost() float64 {
	return s.Beverage.Cost() + .15
}

func (s *Soy) GetDescription() string {
	return s.Beverage.GetDescription() + ", Soy"
}
