package main

import (
	"designpatterns/behavioral/template/barista/concrete"
	"fmt"
)

func main() {
	fmt.Println("=== Template Method Pattern Demo - Barista ===")

	// Crear instancias de diferentes bebidas
	coffee := concrete.NewCoffee()
	tea := concrete.NewTea()
	hotChocolate := concrete.NewHotChocolate()

	fmt.Println("\n=== Preparando Café ===")
	coffee.PrepareRecipe(coffee)

	fmt.Println("\n=== Preparando Té ===")
	tea.PrepareRecipe(tea)

	fmt.Println("\n=== Preparando Chocolate Caliente ===")
	hotChocolate.PrepareRecipe(hotChocolate)

	fmt.Println("\n=== Demo completado ===")
}
