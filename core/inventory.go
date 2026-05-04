package core

import (
	"fmt"
)

type Inventory struct {
	items []InventoryItem
}

func (i *Inventory) Items() []InventoryItem {
	return i.items
}

type InventoryItem interface {
	fmt.Stringer
	Weight() float64
	Value() int
}
