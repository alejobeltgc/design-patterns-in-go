package creator

import "designpatterns/creational/factory/factory_method/product"

type PizzaStore interface {
	OrderPizza(pizzaType string) product.IPizza
	CreatePizza(pizzaType string) product.IPizza
}

type BasePizzaStore struct{}

func (ps *BasePizzaStore) OrderPizza(creator PizzaStore, pizzaType string) product.IPizza {
	pizza := creator.CreatePizza(pizzaType)

	if pizza != nil {
		pizza.Prepare()
		pizza.Bake()
		pizza.Cut()
		pizza.Box()
	}

	return pizza
}
