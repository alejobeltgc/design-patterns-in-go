package iterator

import (
	"designpatterns/behavioral/iterator/dinermerge/models"
	"errors"
)

type DinerMenuIterator struct {
	menuItems []*models.MenuItem
	position  int
}

func NewDinerMenuIterator(items []*models.MenuItem) *DinerMenuIterator {
	return &DinerMenuIterator{
		menuItems: items,
		position:  0,
	}
}

func (d *DinerMenuIterator) HasNext() bool {
	return d.position < len(d.menuItems)
}

func (d *DinerMenuIterator) Next() (*models.MenuItem, error) {
	if !d.HasNext() {
		return nil, errors.New("no more items")
	}
	menuItem := d.menuItems[d.position]
	d.position++
	return menuItem, nil
}
