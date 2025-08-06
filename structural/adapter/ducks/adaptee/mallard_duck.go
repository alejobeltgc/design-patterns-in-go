package adaptee

import "fmt"

type MallardDuck struct{}

func (m *MallardDuck) Quack() {
	fmt.Println("Quack")
}

func (m *MallardDuck) Fly() {
	fmt.Println("Estoy volando!!")
}
