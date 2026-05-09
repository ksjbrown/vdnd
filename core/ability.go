package core

import "fmt"

type Ability int

func (a Ability) BonusModifier() int {
	// offset from 10, divided by 2, rounded down
	offset := int(a) - 10
	if offset >= 0 {
		return offset / 2
	}
	return (offset - 1) / 2
}

// Abilities holds the six ability scores for a character.
type Abilities struct {
	Strength     Ability
	Dexterity    Ability
	Constitution Ability
	Wisdom       Ability
	Intelligence Ability
	Charisma     Ability
}

func (a *Abilities) GetAbility(kind AbilityKind) Ability {
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

	default:
		msg := "unknown ability kind: %v"
		panic(fmt.Sprintf(msg, kind))
	}
}

type AbilityKind int

const (
	AbilityKindStrength     AbilityKind = 0x1
	AbilityKindDexterity    AbilityKind = 0x2
	AbilityKindConstitution AbilityKind = 0x3
	AbilityKindWisdom       AbilityKind = 0x4
	AbilityKindIntelligence AbilityKind = 0x5
	AbilityKindCharisma     AbilityKind = 0x6
)

