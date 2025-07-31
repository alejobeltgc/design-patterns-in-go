package concreteproducts

import (
	"designpatterns/creational/factory/factory_method/product"
	"fmt"
)

type ChicagoStylePepperoniPizza struct {
	product.Pizza
}

func NewChicagoStylePepperoniPizza() *ChicagoStylePepperoniPizza {
	return &ChicagoStylePepperoniPizza{
		Pizza: product.Pizza{
			Name:     "Pizza de pepperoni Chicago Deep Dish",
			Dough:    "Gruesa",
			Sauce:    "Tomate natural",
			Toppings: []string{"Mozzarella extra", "Pepperoni grueso"},
		},
	}
}

func (c *ChicagoStylePepperoniPizza) Prepare() {
	fmt.Printf("Preparando %s con masa %s y salsa %s \n", c.Name, c.Dough, c.Sauce)
	fmt.Printf("Agregando ingredientes: %v\n", c.Toppings)
}

func (c *ChicagoStylePepperoniPizza) Bake() {
	fmt.Println("50 minutos en el horno a 325° (pepperoni deep dish requiere más tiempo)")
}

func (c *ChicagoStylePepperoniPizza) Cut() {
	fmt.Println("Corte en cuadros grandes (estilo Chicago)")
}

func (c *ChicagoStylePepperoniPizza) Box() {
	fmt.Println("Empacado en caja extra profunda para deep dish")
}
