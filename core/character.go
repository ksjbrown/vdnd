package core

// complex
type Character struct {
	ID             CharacterID
	BaseAbilities  Abilities
	BaseArmorClass ArmorClass
	Background     Background
	Class          Class
	Subclasses     []Class
	Origin         Origin
	Species        Species
	Progression    Progression
	Inventory      Inventory
	Loadout        Loadout
	CharacterMutators
}

func (c *Character) GetArmorClass() ArmorClass {
	ac := c.BaseArmorClass
	delta := c.ArmorClassMutator.GetDelta()
	return ac + ArmorClass(delta)
}

func (c *Character) GetAbility(kind AbilityKind) Ability {
	ability := c.BaseAbilities.GetAbility(kind)
	delta := c.AbilitiesMutator.GetDelta(kind)
	return ability + Ability(delta)
}

func (c *Character) GetAttackRoll() int {
	baseRoll := NewDieD20().Roll()
	delta := c.AttackRollMutator.GetDelta()
	return baseRoll + delta
}

type CharacterID uint
