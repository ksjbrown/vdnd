package core

import "fmt"

type SupportsAbilties interface {
	GetAbility(AbilityKind) *Ability
	GetStrength() *Ability
	GetDexterity() *Ability
	GetConstitution() *Ability
	GetIntelligence() *Ability
	GetWisdom() *Ability
	GetCharisma() *Ability
}

type AbilityData struct {
	Strength     int
	Dexterity    int
	Constitution int
	Wisdom       int
	Intelligence int
	Charisma     int 
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

func (a *Abilities) GetAbility(kind AbilityKind) *Ability {
	switch kind {

	case AbilityKindStrength:
		return &a.Strength

	case AbilityKindDexterity:
		return &a.Dexterity

	case AbilityKindConstitution:
		return &a.Constitution

	case AbilityKindWisdom:
		return &a.Wisdom

	case AbilityKindIntelligence:
		return &a.Intelligence

	case AbilityKindCharisma:
		return &a.Charisma

	default:
		msg := "unknown ability kind: %v"
		panic(fmt.Sprintf(msg, kind))
	}
}

func (a *Abilities) GetStrength() *Ability {
	return &a.Strength
}

func (a *Abilities) GetDexterity() *Ability {
	return &a.Dexterity
}

func (a *Abilities) GetConstitution() *Ability {
	return &a.Constitution
}

func (a *Abilities) GetIntelligence() *Ability {
	return &a.Intelligence
}

func (a *Abilities) GetWisdom() *Ability {
	return &a.Wisdom
}

func (a *Abilities) GetCharisma() *Ability {
	return &a.Charisma
}

// Ability represents a single Ability score (STR, DEX, etc.)
//
// An Ability Score consists of a BaseScore, as well as any number of Bonuses, like:
//
// - Character Creation Ability Point distributions
// - Feats
// - Temporary Buffs or Debuffs
type Ability struct {
	SupportsBonuses
	BaseScore int
}

func NewAbility(baseScore int) *Ability {
	return &Ability{
		SupportsBonuses: NewBonusMap(),
		BaseScore:       baseScore,
	}
}

func (a *Ability) Score() int {
	return a.BaseScore + a.GetBonusTotalValue()
}

func (a *Ability) Modifier() int {
	// score offset from 10, divided by 2, rounded down
	offset := a.Score() - 10
	if offset >= 0 {
		return offset / 2
	}
	return (offset - 1) / 2
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
