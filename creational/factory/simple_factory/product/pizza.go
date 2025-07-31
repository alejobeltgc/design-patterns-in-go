package product

import "fmt"

type IPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type Pizza struct {
	Name  string
	Dough string
	Sauce string
}

func (p *Pizza) DefaultBox() {
	fmt.Println("Empacado en caja estándar de la pizzería")
}
