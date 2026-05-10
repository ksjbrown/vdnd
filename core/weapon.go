package core

import "encoding/json"

type Weapon struct {
	ID              WeaponID
	Name 			string
	Category        WeaponCategory
	Properties      []WeaponProperty
	MasteryProperty WeaponMasteryProperty
}

type WeaponID uint

type WeaponCategory int

const (
	WeaponCategorySimple  WeaponCategory = 0x1
	WeaponCategoryMartial WeaponCategory = 0x2
)

type WeaponPropertyData struct {
	Kind WeaponPropertyKind
	Data json.RawMessage
}

type WeaponProperty interface {
	GetWeaponPropertyKind() WeaponPropertyKind
	// supply modifiers, if they exist? Attack Roll modifiers, based on Character stats
}

type WeaponPropertyKind int

const (
	WeaponPropertyKindAmmunition WeaponPropertyKind = 0x01
	WeaponPropertyKindFinesse    WeaponPropertyKind = 0x02
	WeaponPropertyKindHeavy      WeaponPropertyKind = 0x03
	WeaponPropertyKindLight      WeaponPropertyKind = 0x04
	WeaponPropertyKindLoading    WeaponPropertyKind = 0x05
	WeaponPropertyKindReach      WeaponPropertyKind = 0x06
	WeaponPropertyKindThrown     WeaponPropertyKind = 0x07
	WeaponPropertyKindTwoHanded  WeaponPropertyKind = 0x08
	WeaponPropertyKindVersatile  WeaponPropertyKind = 0x09
)

type WeaponMasteryProperty interface {
	GetWeaponMasteryPropertyKind() WeaponMasteryPropertyKind
}

type WeaponMasteryPropertyKind int

const (
	WeaponMasteryPropertyKindCleave WeaponMasteryPropertyKind = 0x01
	WeaponMasteryPropertyKindGraze  WeaponMasteryPropertyKind = 0x02
	WeaponMasteryPropertyKindNick   WeaponMasteryPropertyKind = 0x03
	WeaponMasteryPropertyKindPush   WeaponMasteryPropertyKind = 0x04
	WeaponMasteryPropertyKindSap    WeaponMasteryPropertyKind = 0x05
	WeaponMasteryPropertyKindSlow   WeaponMasteryPropertyKind = 0x06
	WeaponMasteryPropertyKindTopple WeaponMasteryPropertyKind = 0x07
	WeaponMasteryPropertyKindVex    WeaponMasteryPropertyKind = 0x08
)
