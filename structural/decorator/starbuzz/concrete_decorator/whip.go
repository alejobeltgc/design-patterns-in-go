package concretedecorator

import "designpatterns/structural/decorator/starbuzz/component"

type Whip struct {
	Beverage component.Beverage
}

func NewWhip(b component.Beverage) *Whip {
	if b == nil {
		panic("beverage cannot be nil")
	}
	return &Whip{
		Beverage: b,
	}
}

func (w *Whip) Cost() float64 {
	return w.Beverage.Cost() + .10
}

func (w *Whip) GetDescription() string {
	return w.Beverage.GetDescription() + ", Whip"
}
