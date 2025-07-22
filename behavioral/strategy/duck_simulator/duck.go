package main

import (
	"designpatterns/behavioral/strategy/duck_simulator/fly_behavior"
	"designpatterns/behavioral/strategy/duck_simulator/quack_behavior"
	"fmt"
)

type Duck struct {
	flyBehavior   fly_behavior.FlyBehavior
	quackBehavior quack_behavior.QuackBehavior
}

func (d *Duck) PerformFly() {
	if d.flyBehavior != nil {
		d.flyBehavior.Fly()
	}
}

func (d *Duck) PerformQuack() {
	if d.quackBehavior != nil {
		d.quackBehavior.Quack()
	}
}

func (d *Duck) SetFlyBehavior(fb fly_behavior.FlyBehavior) {
	d.flyBehavior = fb
}

func (d *Duck) SetQuackBehavior(qb quack_behavior.QuackBehavior) {
	d.quackBehavior = qb
}

func (d *Duck) Display() {
	fmt.Println("I'm generic Duck!!")
}

func (d *Duck) Swim() {
	fmt.Println("All ducks float, even decoys!")
}
