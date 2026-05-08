package core

type Weapon struct {
	ID              WeaponID
	Name 			string
	Category        WeaponCategory
	Properties      []WeaponProperty
	MasteryProperty WeaponMasteryProperty

	equipCondition func(*Character) bool
}

func (w *Weapon) IsEquippableBy(c *Character) bool {
	if w.equipCondition == nil {
		return true
	}
	return w.equipCondition(c)
}


type WeaponID uint

type WeaponCategory int

const (
	WeaponCategorySimple  WeaponCategory = 0x1
	WeaponCategoryMartial WeaponCategory = 0x2
)


