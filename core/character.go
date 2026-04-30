package core

// TODO: merge all character related data into this type.

// Character represents any person, creature, or actor in the game, that can perform actions, be interacted with etc.
//
// Should contain all data required for character functionality.
type Character struct {
	id CharacterID
	Entity
	Background  Background
	Class       Class
	Subclasses  []Class
	Origin      Origin
	Species     Species
	AC          ArmorClass
	Progression Progression
	Inventory   Inventory
	Loadout     Loadout
}

func (c *Character) ID() CharacterID {
	return c.id
}



type CharacterID uint
