package core

// Combat represents the state of a given combat encounter
type Combat struct{
	ID CombatID
	Combatants []Combatant
}

type CombatID uint64

type Combatant interface{}
