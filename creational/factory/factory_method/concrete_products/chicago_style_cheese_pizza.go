package concreteproducts

import (
	"designpatterns/creational/factory/factory_method/product"
	"fmt"
)

type ChicagoStyleCheesePizza struct {
	product.Pizza
}

func NewChicagoStyleCheesePizza() *ChicagoStyleCheesePizza {
	return &ChicagoStyleCheesePizza{
		Pizza: product.Pizza{
			Name:     "Pizza de quesos Chicago Deep Dish",
			Dough:    "Gruesa",
			Sauce:    "Tomate natural",
			Toppings: []string{"Mozzarella extra", "Parmesano"},
		},
	}
}

func (c *ChicagoStyleCheesePizza) Prepare() {
	fmt.Printf("Preparando %s con masa %s y salsa %s \n", c.Name, c.Dough, c.Sauce)
	fmt.Printf("Agregando ingredientes: %v\n", c.Toppings)
}

func (c *ChicagoStyleCheesePizza) Bake() {
	fmt.Println("45 minutos en el horno a 325° (deep dish requiere más tiempo)")
}

func (c *ChicagoStyleCheesePizza) Cut() {
	fmt.Println("Corte en cuadros (estilo Chicago)")
}

func (c *ChicagoStyleCheesePizza) Box() {
	fmt.Println("Empacado en caja extra profunda para deep dish")
}
