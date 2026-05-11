package core

// SupportsArmorClass represents a type that can provide an armor class value
//
// Usually this means that anything that can attack, can attack this thing.
type SupportsArmorClass interface {
	SupportsBonuses
	SupportsEffectsArmorClass
	GetArmorClass() int
}

type SupportsEffectsArmorClass interface {
	AddArmorClassEffect(Effect[SupportsArmorClass])
	RemoveArmorClassEffect(Effect[SupportsArmorClass])
}

// ArmorClass provides methods to represent something that can provide an ArmorClass value.
type ArmorClass struct {
	SupportsBonuses
	BaseValue int
	BaseValueBonuses SupportsBonuses

	effects *EffectMap[SupportsArmorClass]
}

func NewArmorClass(baseValue int) *ArmorClass {
	ac := &ArmorClass{
		SupportsBonuses: NewBonusMap(),
		BaseValue: baseValue,
		BaseValueBonuses: NewBonusMap(),
	}
	ac.effects = NewEffectMap(SupportsArmorClass(ac))
	return ac
}

func (ac *ArmorClass) GetArmorClass() int {
	// only one base AC bonus is allowed,
	baseValueBonus := 0
	for bonus := range ac.BaseValueBonuses.IterateBonuses() {
		baseValueBonus = max(baseValueBonus, bonus.GetBonusValue())
	}
	return ac.BaseValue + baseValueBonus + ac.GetCombinedBonusValue()
}

func (ac *ArmorClass) AddArmorClassEffect(effect Effect[SupportsArmorClass]) {
	ac.effects.AddEffect(effect)
}

func (ac *ArmorClass) RemoveArmorClassEffect(effect Effect[SupportsArmorClass]) {
	ac.effects.RemoveEffect(effect)
}
