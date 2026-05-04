package core

type Item struct {
	weight int
	value int
}

func (i *Item) Weight() int {
	return i.weight
}

func (i *Item) Value() int {
	return i.value
}
