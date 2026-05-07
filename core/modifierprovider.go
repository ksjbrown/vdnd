package core

// ModifierProvider represents any type that provides a method ModifierBonus()
type ModifierProvider interface {
	Modifier() int
}

