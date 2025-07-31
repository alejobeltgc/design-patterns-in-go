package factory

import (
	concreteproducts "designpatterns/creational/factory/simple_factory/concrete_products"
	"designpatterns/creational/factory/simple_factory/product"
	"fmt"
)

type PizzaFactory struct{}

func NewPizzaFactory() *PizzaFactory {
	return &PizzaFactory{}
}

func (f *PizzaFactory) CreatePizza(pizzaType string) product.IPizza {
	switch pizzaType {
	case "cheese":
		return concreteproducts.NewCheesePizza()
	case "pepperoni":
		return concreteproducts.NewPepperoniPizza()
	case "veggie":
		return concreteproducts.NewVeggiePizza()
	default:
		fmt.Printf("No tenemos disponible: %s\n", pizzaType)
		return nil
	}
}
