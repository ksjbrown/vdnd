package core

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

	modifiers CharacterModifiers
}

func (c *Character) GetAbility(kind AbilityKind) Ability {
	ability := c.BaseAbilities.GetAbility(kind)
	bonus := c.modifiers.Abilities.GetBonus(kind)
	return ability + Ability(bonus)
}

func (c *Character) GetArmorClass() ArmorClass {
	ac := c.BaseArmorClass
	bonus := c.modifiers.ArmorClass.GetBonus()
	return ac + ArmorClass(bonus)
}

func (c *Character) GetAttackRollResult(target ArmorClass) *D20TestResult {
	die := NewDieD20()
	modifiers := &c.modifiers.AttackRolls
	bonus := modifiers.GetBonus()
	advantage := modifiers.GetAdvantage()
	test := NewD20Test(die, bonus, advantage, int(target))
	return test.DoTest()
}

type CharacterID uint


