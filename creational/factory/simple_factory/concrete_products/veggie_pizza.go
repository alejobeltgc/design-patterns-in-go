package concreteproducts

import (
	"designpatterns/creational/factory/simple_factory/product"
	"fmt"
)

type VeggiePizza struct {
	product.Pizza
}

func NewVeggiePizza() *VeggiePizza {
	return &VeggiePizza{
		Pizza: product.Pizza{
			Name:  "Pizza vegetariana",
			Dough: "Americana",
			Sauce: "Española",
		},
	}
}

func (c *VeggiePizza) Prepare() {
	fmt.Printf("Preparando %s con masa %s y salsa %s \n", c.Name, c.Dough, c.Sauce)
}

func (c *VeggiePizza) Bake() {
	fmt.Println("20 minutos en el horno a 325° (temperatura baja para vegetales)")
}

func (c *VeggiePizza) Cut() {
	fmt.Println("Corte en 8 cuadros para mejor presentación")
}

func (c *VeggiePizza) Box() {
	fmt.Println("Empacado en caja eco-friendly biodegradable")
}
