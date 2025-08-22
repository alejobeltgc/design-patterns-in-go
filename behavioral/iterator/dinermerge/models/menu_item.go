package models

type MenuItem struct {
	Name        string
	Description string
	Vegetarian  bool
	Price       float64
}

func NewMenuItem(name, description string, vegetarian bool, price float64) *MenuItem {
	return &MenuItem{
		Name:        name,
		Description: description,
		Vegetarian:  vegetarian,
		Price:       price,
	}
}

func (m *MenuItem) GetName() string {
	return m.Name
}

func (m *MenuItem) GetDescription() string {
	return m.Description
}

func (m *MenuItem) IsVegetarian() bool {
	return m.Vegetarian
}

func (m *MenuItem) GetPrice() float64 {
	return m.Price
}
