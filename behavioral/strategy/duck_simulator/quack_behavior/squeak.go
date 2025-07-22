package quack_behavior

import "fmt"

type Squeak struct{}

func (q *Squeak) Quack() {
	fmt.Println("Squeak")
}
