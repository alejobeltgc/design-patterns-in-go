package adapter

import (
	"designpatterns/structural/adapter/ducks/interfaces"
	"fmt"
)

type TurkeyAdapter struct {
	turkey interfaces.Turkey
}

func NewTurkeyAdapter(turkey interfaces.Turkey) *TurkeyAdapter {
	if turkey == nil {
		panic("turkey cannot be nil")
	}
	return &TurkeyAdapter{
		turkey: turkey,
	}
}

func (ta *TurkeyAdapter) Quack() {
	ta.turkey.Gobble()
}

func (ta *TurkeyAdapter) Fly() {
	fmt.Println("Adaptando vuelo de pavo a vuelo de pato...")
	for i := 0; i < 5; i++ {
		ta.turkey.Fly()
	}
}
