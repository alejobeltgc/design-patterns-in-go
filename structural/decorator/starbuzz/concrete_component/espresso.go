package concretecomponent

type Espresso struct {
	Description string
}

func NewEspresso() *Espresso {
	return &Espresso{
		Description: "Espresso",
	}
}

func (e *Espresso) Cost() float64 {
	return 1.99
}

func (e *Espresso) GetDescription() string {
	return e.Description
}
