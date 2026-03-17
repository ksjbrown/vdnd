package core

type Combat struct {
	Combatants []Combatant
}

type Combatant struct {
	Character
	InitiativeRoll int
	Position Position
}

