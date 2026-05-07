package core

import "fmt"

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
