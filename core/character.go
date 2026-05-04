package core

// TODO: merge all character related data into this type.

// Character represents any person, creature, or actor in the game, that can perform actions, be interacted with etc.
//
// Should contain all data required for character functionality.
type Character struct {
	ID          CharacterID
	background  *Background
	class       *Class
	subclasses  []Class
	origin      *Origin
	species     *Species
	ac          *ArmorClass
	progression *Progression
	inventory   *Inventory
	loadout     *Loadout
}

func (c *Character) Background() *Background {
	return c.background
}

func (c *Character) Class() *Class {
	return c.class
}

func (c *Character) Subclasses() []Class {
	return c.subclasses
}

func (c *Character) Origin() *Origin {
	return c.origin
}

func (c *Character) Species() *Species {
	return c.species
}

func (c *Character) AC() *ArmorClass {
	return c.ac
}

func (c *Character) Progression() *Progression {
	return c.progression
}

func (c *Character) Inventory() *Inventory {
	return c.inventory
}

func (c *Character) Loadout() *Loadout {
	return c.loadout
}
type CharacterID uint
