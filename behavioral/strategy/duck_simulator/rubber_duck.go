package main

import (
	"designpatterns/behavioral/strategy/duck_simulator/fly_behavior"
	"designpatterns/behavioral/strategy/duck_simulator/quack_behavior"
	"fmt"
)

type RubberDuck struct {
	Duck
}

func NewRubberDuck() *RubberDuck {
	rubber := &RubberDuck{}
	rubber.SetFlyBehavior(&fly_behavior.FlyNoWay{})
	rubber.SetQuackBehavior(&quack_behavior.Squeak{})
	return rubber
}

func (r *RubberDuck) Display() {
	fmt.Println("I'm a rubber duckie")
}
