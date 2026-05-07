package core

type ArmorClassProvider interface {
	GetArmorClass() ArmorClass
}

type ArmorClass int
