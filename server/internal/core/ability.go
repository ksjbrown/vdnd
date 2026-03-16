package core

import "math"

type Ability struct {
	Type AbilityType
	Score AbilityScore
}

type AbilityType int

const (
	Strength     AbilityType = 0x1
	Dexterity    AbilityType = 0x2
	Constitution AbilityType = 0x3
	Wisdom       AbilityType = 0x4
	Intelligence AbilityType = 0x5
	Charisma     AbilityType = 0x6
)

// AbilityScore represents the value that an Ability can have.
type AbilityScore int

// Modifier returns the score modifier for the current score value
func (s AbilityScore) Modifier() int {
	return int(math.Floor((float64(s) - 10) / 2))
}

// Abilities represents the base Character ability scores chosen during character creation.
// The actual effective score value may be higher, depending on things like Feats, status effects, etc. 
type Abilities struct {
	Strength     AbilityScore
	Dexterity    AbilityScore
	Constitution AbilityScore
	Wisdom       AbilityScore
	Intelligence AbilityScore
	Charisma     AbilityScore
}
