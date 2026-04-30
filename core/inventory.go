package core

import (
	"fmt"
)

type Inventory struct {
	Items []InventoryItem
}

type InventoryItem interface {
	fmt.Stringer
	Weight() float64
	// Value()
}
