package concrete

import (
	"designpatterns/behavioral/template/barista/template"
	"fmt"
)

type HotChocolate struct {
	template.CaffeineBeverage
}

func NewHotChocolate() *HotChocolate {
	return &HotChocolate{}
}

func (h *HotChocolate) Brew() {
	fmt.Println("Mezclando chocolate en polvo con agua caliente")
}

func (h *HotChocolate) AddCondiments() {
	fmt.Println("Agregando malvaviscos y crema batida")
}

func (h *HotChocolate) CustomerWantsCondiments() bool {
	fmt.Println("¡El chocolate caliente siempre lleva malvaviscos!")
	return true
}
