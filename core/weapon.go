package core

type Weapon struct {
	Item
	ID WeaponID
	kind WeaponKind
	properties []WeaponProperty
}

type WeaponID uint

type WeaponKind int

type WeaponProperty struct {
	kind WeaponPropertyKind
}


type WeaponPropertyKind int

