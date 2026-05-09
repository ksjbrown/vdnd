package core

type Character struct {
	ID             CharacterID
	Abilities
	ArmorClass
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

type CharacterID uint
