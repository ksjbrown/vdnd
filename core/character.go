package core

type CharacterData struct {
	AbilitiesData
}

type Character struct {
	ID             CharacterID
	SupportsAbilities
	SupportsArmorClass
	SupportsFeats
	Background     Background
	Class          Class
	Subclasses     []Class
	Origin         Origin
	Species        Species
	Progression    Progression
	Inventory      Inventory
	Loadout        Loadout
}

func NewCharacter(d *CharacterData) *Character {
	return &Character{
		SupportsAbilities: &Abilities{},
		// TODO: complete
	}
}


type CharacterID uint
