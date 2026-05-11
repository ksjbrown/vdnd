package core

type CharacterData struct {

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

	modifiers CharacterModifiers
}

func NewCharacter(d *CharacterData) *Character {
	return &Character{
		SupportsAbilities: &Abilities{},
	}
}


type CharacterID uint
