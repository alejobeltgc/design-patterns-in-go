package concreteproducts

import (
	"designpatterns/creational/factory/simple_factory/product"
	"fmt"
)

type CheesePizza struct {
	product.Pizza
}

func NewCheesePizza() *CheesePizza {
	return &CheesePizza{
		Pizza: product.Pizza{
			Name:  "Pizza de quesos",
			Dough: "Tradicional",
			Sauce: "Carbonara",
		},
	}
}

func (c *CheesePizza) Prepare() {
	fmt.Printf("Preparando %s con masa %s y salsa %s \n", c.Name, c.Dough, c.Sauce)
}

func (c *CheesePizza) Bake() {
	fmt.Println("25 minutos en el horno a 350°")
}

func (c *CheesePizza) Cut() {
	fmt.Println("Corte en 8 triángulos tradicionales")
}

func (c *CheesePizza) Box() {
	c.DefaultBox()
}
