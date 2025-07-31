package concreteproducts

import (
	"designpatterns/creational/factory/factory_method/product"
	"fmt"
)

type NyStylePepperoniPizza struct {
	product.Pizza
}

func NewNyStylePepperoniPizza() *NyStylePepperoniPizza {
	return &NyStylePepperoniPizza{
		Pizza: product.Pizza{
			Name:     "Pizza de pepperoni NY",
			Dough:    "Delgada",
			Sauce:    "Marinara",
			Toppings: []string{"Mozzarella", "Pepperoni"},
		},
	}
}

func (c *NyStylePepperoniPizza) Prepare() {
	fmt.Printf("Preparando %s con masa %s y salsa %s \n", c.Name, c.Dough, c.Sauce)
	fmt.Printf("Agregando ingredientes: %v\n", c.Toppings)
}

func (c *NyStylePepperoniPizza) Bake() {
	fmt.Println("30 minutos en el horno a 375°")
}

func (c *NyStylePepperoniPizza) Cut() {
	fmt.Println("Corte en 8 triángulos diagonales")
}

func (c *NyStylePepperoniPizza) Box() {
	fmt.Println("Empacado en caja oficial NY Pizza")
}
