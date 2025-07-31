package product

type IPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type Pizza struct {
	Name     string
	Dough    string
	Sauce    string
	Toppings []string
}
