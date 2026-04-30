package core

type Weapon struct {
	Item
	ID WeaponID
	Type WeaponType
	Properties []WeaponProperty
}

type WeaponID uint

type WeaponType int

type WeaponProperty struct {
	Type WeaponPropertyType
	Args WeaponPropertyArgs // some json representation of additional data, e.g range of weapon
}

type WeaponPropertyArgs any

type WeaponPropertyType int

