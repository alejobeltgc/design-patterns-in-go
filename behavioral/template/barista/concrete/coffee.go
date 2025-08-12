package concrete

import (
	"designpatterns/behavioral/template/barista/template"
	"fmt"
	"strings"
)

type Coffee struct {
	template.CaffeineBeverage
}

func NewCoffee() *Coffee {
	return &Coffee{}
}

func (c *Coffee) Brew() {
	fmt.Println("Filtrando el café a través del filtro")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("Agregando azúcar y leche")
}

func (c *Coffee) CustomerWantsCondiments() bool {
	answer := c.getUserInput()
	return strings.ToLower(answer) == "y"
}

func (c *Coffee) getUserInput() string {
	fmt.Print("¿Te gustaría leche y azúcar con tu café (y/n)? ")

	answer := "y"
	fmt.Println(answer)
	return answer
}
