package concretecreator

import (
	concreteproducts "designpatterns/creational/factory/factory_method/concrete_products"
	"designpatterns/creational/factory/factory_method/creator"
	"designpatterns/creational/factory/factory_method/product"
	"fmt"
)

type NyStylePizzaStore struct {
	creator.BasePizzaStore
}

func (ps *NyStylePizzaStore) OrderPizza(pizzaType string) product.IPizza {
	return ps.BasePizzaStore.OrderPizza(ps, pizzaType)
}

func (ps *NyStylePizzaStore) CreatePizza(pizzaType string) product.IPizza {
	switch pizzaType {
	case "cheese":
		return concreteproducts.NewNyStyleCheesePizza()
	case "pepperoni":
		return concreteproducts.NewNyStylePepperoniPizza()
	default:
		fmt.Printf("No tenemos disponible pizza %s estilo NY\n", pizzaType)
		return nil
	}
}
