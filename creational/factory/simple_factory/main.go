package main

import (
	"designpatterns/creational/factory/simple_factory/factory"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Simple Factory Pattern Demo ===")

	pizzaFactory := factory.NewPizzaFactory()

	pizzaTypes := []string{"cheese", "pepperoni", "veggie", "hawaiana"}

	for _, pizzaType := range pizzaTypes {
		fmt.Printf("\n--- Ordenando pizza %s ---\n", pizzaType)
		fmt.Println(strings.Repeat("-", 35))

		pizza := pizzaFactory.CreatePizza(pizzaType)

		if pizza != nil {
			pizza.Prepare()
			pizza.Bake()
			pizza.Cut()
			pizza.Box()
			fmt.Println("¡Pizza lista para entregar!")
		} else {
			fmt.Println("Lo sentimos, no pudimos hacer esa pizza.")
		}
	}

	fmt.Println("\n=== Demo completado ===")
}
