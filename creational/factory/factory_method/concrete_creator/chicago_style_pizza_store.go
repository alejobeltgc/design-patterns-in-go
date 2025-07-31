package concretecreator

import (
	concreteproducts "designpatterns/creational/factory/factory_method/concrete_products"
	"designpatterns/creational/factory/factory_method/creator"
	"designpatterns/creational/factory/factory_method/product"
	"fmt"
)

type ChicagoStylePizzaStore struct {
	creator.BasePizzaStore
}

func (ps *ChicagoStylePizzaStore) OrderPizza(pizzaType string) product.IPizza {
	return ps.BasePizzaStore.OrderPizza(ps, pizzaType)
}

func (ps *ChicagoStylePizzaStore) CreatePizza(pizzaType string) product.IPizza {
	switch pizzaType {
	case "cheese":
		return concreteproducts.NewChicagoStyleCheesePizza()
	case "pepperoni":
		return concreteproducts.NewChicagoStylePepperoniPizza()
	default:
		fmt.Printf("No tenemos disponible pizza %s estilo Chicago\n", pizzaType)
		return nil
	}
}
