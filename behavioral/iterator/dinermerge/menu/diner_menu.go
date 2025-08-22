package menu

import (
	"designpatterns/behavioral/iterator/dinermerge/iterator"
	"designpatterns/behavioral/iterator/dinermerge/models"
)

type DinerMenu struct {
	menuItems []*models.MenuItem
	maxItems  int
	itemCount int
}

func NewDinerMenu() *DinerMenu {
	menuItems := make([]*models.MenuItem, 6)
	menuItems[0] = models.NewMenuItem("Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99)
	menuItems[1] = models.NewMenuItem("BLT", "Bacon with lettuce & tomato on whole wheat", false, 2.99)
	menuItems[2] = models.NewMenuItem("Soup of the day", "Soup of the day, with a side of potato salad", false, 3.29)
	menuItems[3] = models.NewMenuItem("Hotdog", "A hot dog, with sauerkraut, relish, onions, topped with cheese", false, 3.05)
	menuItems[4] = models.NewMenuItem("Steamed Veggies and Brown Rice", "Steamed vegetables over brown rice", true, 3.99)
	menuItems[5] = models.NewMenuItem("Pasta", "Spaghetti with Marinara Sauce, and a slice of sourdough bread", true, 3.89)

	return &DinerMenu{
		menuItems: menuItems,
		maxItems:  6,
		itemCount: 0,
	}
}

func (d *DinerMenu) AddItem(name, description string, vegetarian bool, price float64) {
	if d.itemCount >= d.maxItems {
		return
	}
	item := models.NewMenuItem(name, description, vegetarian, price)
	d.menuItems[d.itemCount] = item
	d.itemCount++
}

func (d *DinerMenu) CreateIterator() iterator.Iterator[*models.MenuItem] {
	return iterator.NewDinerMenuIterator(d.menuItems)
}
