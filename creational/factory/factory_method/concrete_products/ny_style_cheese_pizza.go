package concreteproducts

import (
	"designpatterns/creational/factory/factory_method/product"
	"fmt"
)

type NyStyleCheesePizza struct {
	product.Pizza
}

func NewNyStyleCheesePizza() *NyStyleCheesePizza {
	return &NyStyleCheesePizza{
		Pizza: product.Pizza{
			Name:     "Pizza de quesos NY",
			Dough:    "Tradicional",
			Sauce:    "Ranch",
			Toppings: []string{"Queso azul", "Bocadillo"},
		},
	}
}

func (c *NyStyleCheesePizza) Prepare() {
	fmt.Printf("Preparando %s con masa %s y salsa %s \n", c.Name, c.Dough, c.Sauce)
}

func (c *NyStyleCheesePizza) Bake() {
	fmt.Println("25 minutos en el horno a 350°")
}

func (c *NyStyleCheesePizza) Cut() {
	fmt.Println("Corte en 8 triángulos tradicionales")
}

func (c *NyStyleCheesePizza) Box() {
	fmt.Println("Empacado en caja especial resistente al calor")
}
