package menu

import (
	"designpatterns/behavioral/iterator/dinermerge/iterator"
	"designpatterns/behavioral/iterator/dinermerge/models"
)

type PancakeHouseMenu struct {
	menuItems []*models.MenuItem
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	menuItems := []*models.MenuItem{
		models.NewMenuItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs, and toast", true, 2.99),
		models.NewMenuItem("Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", false, 2.99),
		models.NewMenuItem("Blueberry Pancakes", "Pancakes made with fresh blueberries", true, 3.49),
		models.NewMenuItem("Waffles", "Waffles with your choice of blueberries or strawberries", true, 3.59),
	}

	return &PancakeHouseMenu{menuItems: menuItems}
}

func (p *PancakeHouseMenu) AddItem(name, description string, vegetarian bool, price float64) {
	item := models.NewMenuItem(name, description, vegetarian, price)
	p.menuItems = append(p.menuItems, item)
}

func (p *PancakeHouseMenu) CreateIterator() iterator.Iterator[*models.MenuItem] {
	return iterator.NewPancakeHouseMenuIterator(p.menuItems)
}
