package core

import (
	"fmt"
	"math"
)

// Ability represents a single D&D ability score value.
type Ability int

// NewAbility creates a new Ability with the given value.
func NewAbility(value int) Ability {
	// TODO: bounds check
	return Ability(value)
}

// Modifier returns the score modifier for the current score value.
func (s Ability) Modifier() int {
	score := float64(s)
	value := (score - 10) / 2
	return int(math.Floor(value))
}

// AbilityKind identifies a specific type of D&D ability.
type AbilityKind int

const (
	AbilityKindStrength AbilityKind = 0x1
	AbilityKindDexterity AbilityKind = 0x2
	AbilityKindConstitution AbilityKind = 0x3
	AbilityKindWisdom AbilityKind = 0x4
	AbilityKindIntelligence AbilityKind = 0x5
	AbilityKindCharisma AbilityKind = 0x6
)

// Abilities holds the six ability scores for a character.
type Abilities struct {
	Strength     Ability
	Dexterity    Ability
	Constitution Ability
	Wisdom       Ability
	Intelligence Ability
	Charisma     Ability
}

// Get returns the ability score for the specified AbilityKind. It panics if the kind is unknown.
func (a *Abilities) Get(kind AbilityKind) Ability {

	switch kind {

	case AbilityKindStrength:
		return a.Strength

	case AbilityKindDexterity:
		return a.Dexterity

	case AbilityKindConstitution:
		return a.Constitution

	case AbilityKindWisdom:
		return a.Wisdom

	case AbilityKindIntelligence:
		return a.Intelligence

	case AbilityKindCharisma:
		return a.Charisma
	}

	panic(fmt.Sprintf("unknown ability kind: %v", kind))
}
