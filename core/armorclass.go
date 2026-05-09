package core

// SupportsArmorClass represents a type that can provide an armor class value
//
// Usually this means that anything that can attack, can attack this thing.
type SupportsArmorClass interface {
	GetArmorClass() int
}

// ArmorClass provides methods to represent something that can provide an ArmorClass value.
type ArmorClass struct {
	SupportsBonuses
	BaseValue int
	BaseValueBonuses []Bonus
}

func NewArmorClass(baseValue int) *ArmorClass {
	return &ArmorClass{
		SupportsBonuses: NewSimpleBonusSupporter(),
		BaseValue: baseValue,
		BaseValueBonuses: make([]Bonus, 0),
	}
}

func (ac *ArmorClass) GetArmorClass() int {
	baseValueBonus := 0
	for _, bonus := range ac.BaseValueBonuses {
		baseValueBonus = max(baseValueBonus, bonus.GetBonusValue())
	}
	return ac.BaseValue + baseValueBonus + ac.ComputeBonuses()
}
