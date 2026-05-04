package core

// TODO: merge all character related data into this type.

// Character represents any person, creature, or actor in the game, that can perform actions, be interacted with etc.
//
// Should contain all data required for character functionality.
type CharacterData struct {
	ID          CharacterID
	Background  *Background
	Class       *Class
	Subclasses  []Class
	Origin      *Origin
	Species     *Species
	AC          *ArmorClass
	Progression *Progression
	Inventory   *Inventory
	Loadout     *Loadout
}

type Character struct {
	data *CharacterData
}

func NewCharacter(data *CharacterData) *Character {
	return &Character{ data: data }
}

func (c *Character) Background() *Background {
	return c.data.Background
}

func (c *Character) Class() *Class {
	return c.data.Class
}

func (c *Character) Subclasses() []Class {
	return c.data.Subclasses
}

func (c *Character) Origin() *Origin {
	return c.data.Origin
}

func (c *Character) Species() *Species {
	return c.data.Species
}

func (c *Character) AC() *ArmorClass {
	return c.data.AC
}

func (c *Character) Progression() *Progression {
	return c.data.Progression
}

func (c *Character) Inventory() *Inventory {
	return c.data.Inventory
}

func (c *Character) Loadout() *Loadout {
	return c.data.Loadout
}

type CharacterID uint
