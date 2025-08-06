package adapter

import (
	"designpatterns/structural/adapter/ducks/interfaces"
	"fmt"
	"math/rand"
	"time"
)

type DuckAdapter struct {
	duck interfaces.Duck
	rand *rand.Rand
}

func NewDuckAdapter(duck interfaces.Duck) *DuckAdapter {
	if duck == nil {
		panic("duck cannot be nil")
	}
	return &DuckAdapter{
		duck: duck,
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (da *DuckAdapter) Gobble() {
	da.duck.Quack()
}

func (da *DuckAdapter) Fly() {
	if da.rand.Intn(5) == 0 {
		fmt.Println("No puedo volar tan lejos como un pato normal...")
	} else {
		da.duck.Fly()
	}
}
