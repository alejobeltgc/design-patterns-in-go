package concretecomponent

type Decaf struct {
	Description string
}

func NewDecaf() *Decaf {
	return &Decaf{
		Description: "Decaf",
	}
}

func (e *Decaf) Cost() float64 {
	return 1.05
}

func (d *Decaf) GetDescription() string {
	return d.Description
}
