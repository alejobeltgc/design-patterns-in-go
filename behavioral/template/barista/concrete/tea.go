package concrete

import (
	"designpatterns/behavioral/template/barista/template"
	"fmt"
	"strings"
)

type Tea struct {
	template.CaffeineBeverage
}

func NewTea() *Tea {
	return &Tea{}
}

func (t *Tea) Brew() {
	fmt.Println("Infusionando el té en agua caliente")
}

func (t *Tea) AddCondiments() {
	fmt.Println("Agregando limón")
}

func (t *Tea) CustomerWantsCondiments() bool {
	answer := t.getUserInput()
	return strings.ToLower(answer) == "y"
}

func (t *Tea) getUserInput() string {
	fmt.Print("¿Te gustaría limón con tu té (y/n)? ")

	answer := "n"
	fmt.Println(answer)
	return answer
}
