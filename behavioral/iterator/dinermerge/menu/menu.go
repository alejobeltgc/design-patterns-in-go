package menu

import (
	"designpatterns/behavioral/iterator/dinermerge/iterator"
	"designpatterns/behavioral/iterator/dinermerge/models"
)

type Menu interface {
	CreateIterator() iterator.Iterator[*models.MenuItem]
}
