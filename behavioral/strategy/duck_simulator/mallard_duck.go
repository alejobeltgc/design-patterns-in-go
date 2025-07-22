package main

import (
	"designpatterns/behavioral/strategy/duck_simulator/fly_behavior"
	"designpatterns/behavioral/strategy/duck_simulator/quack_behavior"
	"fmt"
)

type MallardDuck struct {
	Duck
}

func NewMallardDuck() *MallardDuck {
	mallard := &MallardDuck{}
	mallard.SetFlyBehavior(&fly_behavior.FlyWithWings{})
	mallard.SetQuackBehavior(&quack_behavior.Quack{})
	return mallard
}

func (m *MallardDuck) Display() {
	fmt.Println("I'm a real Mallard duck")
}
