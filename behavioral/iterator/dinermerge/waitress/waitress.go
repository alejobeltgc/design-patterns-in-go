package waitress

import (
	"designpatterns/behavioral/iterator/dinermerge/iterator"
	"designpatterns/behavioral/iterator/dinermerge/menu"
	"designpatterns/behavioral/iterator/dinermerge/models"
)

type Waitress struct {
	PancakeHouseMenu menu.Menu
	DinerMenu        menu.Menu
}

func NewWaitress(pancakeHouseMenu menu.Menu, dinerMenu menu.Menu) *Waitress {
	return &Waitress{
		PancakeHouseMenu: pancakeHouseMenu,
		DinerMenu:        dinerMenu,
	}
}

func (w *Waitress) PrintMenu() {
	pancakeIterator := w.PancakeHouseMenu.CreateIterator()
	dinerIterator := w.DinerMenu.CreateIterator()

	println("MENU\n----\nBREAKFAST")
	w.printMenu(pancakeIterator)
	println("\nLUNCH")
	w.printMenu(dinerIterator)
}

func (w *Waitress) printMenu(iterator iterator.Iterator[*models.MenuItem]) {
	for iterator.HasNext() {
		menuItem, err := iterator.Next()
		if err != nil {
			println("Error:", err.Error())
			continue
		}
		println(menuItem.Name, "-", menuItem.Description, "(", menuItem.Price, ")")
		if menuItem.Vegetarian {
			println("(Vegetarian)")
		}
	}
}
