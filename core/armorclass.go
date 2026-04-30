package core

type ArmorClass struct {
	base int
}

func (a *ArmorClass) Base() int {
	return a.base
}
