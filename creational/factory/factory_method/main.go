package main

import (
	concretecreator "designpatterns/creational/factory/factory_method/concrete_creator"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Factory Method Pattern Demo ===")

	// Crear las tiendas (concrete creators)
	nyStore := &concretecreator.NyStylePizzaStore{}
	chicagoStore := &concretecreator.ChicagoStylePizzaStore{}

	fmt.Println("\n--- Pedido en tienda de NY ---")
	fmt.Println(strings.Repeat("-", 40))
	pizza1 := nyStore.OrderPizza("cheese")
	fmt.Printf("Ethan ordenó: %s\n", pizza1)

	fmt.Println("\n--- Pedido en tienda de Chicago ---")
	fmt.Println(strings.Repeat("-", 40))
	pizza2 := chicagoStore.OrderPizza("cheese")
	fmt.Printf("Joel ordenó: %s\n", pizza2)

	fmt.Println("\n--- Comparando pepperoni de ambas tiendas ---")
	fmt.Println(strings.Repeat("-", 50))

	fmt.Println("\nNY Pepperoni:")
	nyStore.OrderPizza("pepperoni")

	fmt.Println("\nChicago Pepperoni:")
	chicagoStore.OrderPizza("pepperoni")

	fmt.Println("\n--- Probando pizza no disponible ---")
	fmt.Println(strings.Repeat("-", 40))
	invalidPizza := nyStore.OrderPizza("hawaiana")
	if invalidPizza == nil {
		fmt.Println("No se pudo crear la pizza solicitada")
	}

	fmt.Println("\n=== Demo completado ===")
}
