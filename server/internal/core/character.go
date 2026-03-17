package core

// Character represents any person, creature, or actor in the game, that can perform actions, be interacted with etc.
//
// Should contain all data required for character functionality.
type Character struct {
	ID         CharacterID
	Name       string
	Bio        CharacterBio
	Level      CharacterLevel
	Background Background
	Class      Class
	Subclasses []Class
	Origin     Origin
	Species    Species
	Abilities  Abilities
	HP         Health
	AC         ArmorClass
	Inventory  Inventory
	Loadout    Loadout
}

type CharacterID uint

type CharacterLevel int

type CharacterBio struct {
}
