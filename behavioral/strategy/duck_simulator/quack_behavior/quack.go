package quack_behavior

import "fmt"

type Quack struct{}

func (q *Quack) Quack() {
	fmt.Println("Quack")
}
