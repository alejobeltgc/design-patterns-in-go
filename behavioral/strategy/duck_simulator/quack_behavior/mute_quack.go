package quack_behavior

import "fmt"

type MuteQuack struct{}

func (q *MuteQuack) Quack() {
	fmt.Println("<< Silence >>")
}
