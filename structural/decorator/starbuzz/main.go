package main

import (
	concretecomponent "designpatterns/structural/decorator/starbuzz/concrete_component"
	concretedecorator "designpatterns/structural/decorator/starbuzz/concrete_decorator"
	"fmt"
)

func main() {
	fmt.Println("=== Starbuzz Coffee - Decorator Pattern Demo ===")

	// Crear un espresso base
	espresso := concretecomponent.NewEspresso()
	fmt.Printf("%s: $%.2f\n", espresso.GetDescription(), espresso.Cost())

	// Agregar leche
	espressoWithMilk := concretedecorator.NewMilk(espresso)
	fmt.Printf("%s: $%.2f\n", espressoWithMilk.GetDescription(), espressoWithMilk.Cost())

	// Agregar mocha y whip
	espressoWithMilkMochaWhip := concretedecorator.NewWhip(
		concretedecorator.NewMocha(espressoWithMilk))
	fmt.Printf("%s: $%.2f\n", espressoWithMilkMochaWhip.GetDescription(), espressoWithMilkMochaWhip.Cost())

	fmt.Println("\n--- Otra bebida ---")

	// House Blend con múltiples decoradores
	houseBlend := concretecomponent.NewHouseBlend()
	complexBeverage := concretedecorator.NewWhip(
		concretedecorator.NewMocha(
			concretedecorator.NewSoy(
				concretedecorator.NewMilk(houseBlend))))

	fmt.Printf("%s: $%.2f\n", complexBeverage.GetDescription(), complexBeverage.Cost())
}
