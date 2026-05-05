package core

type Character struct {
	ID          CharacterID
	Abilities   Abilities
	Background  Background
	Class       Class
	Subclasses  []Class
	Origin      Origin
	Species     Species
	Progression Progression
	Inventory   Inventory
	Loadout     Loadout
}

func (c *Character) BaseAC() ArmorClass {
	// Base AC formula, may be calculated differently, e.g. Armor might improve base AC
	dexterity := c.Abilities.Dexterity
	return ArmorClass(10 + dexterity.Modifier())
}

func (c *Character) AC() ArmorClass {
	// we probably want to store a base value, and calculate ac as we need it.
	// things that can modify AC:
	// - dexterity modifier
	// - Armor equipment,
	return c.BaseAC()
}

type CharacterID uint
