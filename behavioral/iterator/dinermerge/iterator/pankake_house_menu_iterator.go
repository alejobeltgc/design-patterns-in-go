package iterator

import (
	"designpatterns/behavioral/iterator/dinermerge/models"
	"errors"
)

type PancakeHouseMenuIterator struct {
	menuItems []*models.MenuItem
	position  int
}

func NewPancakeHouseMenuIterator(items []*models.MenuItem) *PancakeHouseMenuIterator {
	return &PancakeHouseMenuIterator{
		menuItems: items,
		position:  0,
	}
}

func (p *PancakeHouseMenuIterator) HasNext() bool {
	return p.position < len(p.menuItems)
}

func (p *PancakeHouseMenuIterator) Next() (*models.MenuItem, error) {
	if !p.HasNext() {
		return nil, errors.New("no more items")
	}
	menuItem := p.menuItems[p.position]
	p.position++
	return menuItem, nil
}
