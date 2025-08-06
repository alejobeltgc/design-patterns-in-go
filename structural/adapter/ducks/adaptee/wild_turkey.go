package adaptee

import "fmt"

type WildTurkey struct{}

func (w *WildTurkey) Gobble() {
	fmt.Println("Gobble Gobble")
}

func (w *WildTurkey) Fly() {
	fmt.Println("Estoy volando a corta distancia!!")
}
