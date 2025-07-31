package concreteproducts

import (
	"designpatterns/creational/factory/simple_factory/product"
	"fmt"
)

type PepperoniPizza struct {
	product.Pizza
}

func NewPepperoniPizza() *PepperoniPizza {
	return &PepperoniPizza{
		Pizza: product.Pizza{
			Name:  "Pizza de pepperoni",
			Dough: "Romana",
			Sauce: "Mornay",
		},
	}
}

func (c *PepperoniPizza) Prepare() {
	fmt.Printf("Preparando %s con masa %s y salsa %s \n", c.Name, c.Dough, c.Sauce)
}

func (c *PepperoniPizza) Bake() {
	fmt.Println("35 minutos en el horno a 400° (más caliente por el pepperoni)")
}

func (c *PepperoniPizza) Cut() {
	fmt.Println("Corte en 6 triángulos grandes")
}

func (c *PepperoniPizza) Box() {
	fmt.Println("Empacado en caja especial resistente al calor")
}
