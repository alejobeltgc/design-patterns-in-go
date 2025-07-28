package concretecomponent

type DarkRoast struct {
	Description string
}

func NewDarkRoast() *DarkRoast {
	return &DarkRoast{
		Description: "Dark Roast",
	}
}

func (e *DarkRoast) Cost() float64 {
	return .99
}

func (d *DarkRoast) GetDescription() string {
	return d.Description
}
