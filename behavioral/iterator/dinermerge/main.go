package main

import (
	"designpatterns/behavioral/iterator/dinermerge/menu"
	"designpatterns/behavioral/iterator/dinermerge/waitress"
)

func main() {
	pancakeHouseMenu := menu.NewPancakeHouseMenu()
	dinerMenu := menu.NewDinerMenu()

	waitress := waitress.NewWaitress(pancakeHouseMenu, dinerMenu)

	waitress.PrintMenu()
}
