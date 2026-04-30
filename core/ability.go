package core

import "math"

type Ability struct {
	kind  AbilityKind
	score int
}

func NewAbility(kind AbilityKind, score int) *Ability {
	return &Ability{
		kind:  kind,
		score: score,
	}
}

func (a *Ability) Kind() AbilityKind {
	return a.kind
}

func (a *Ability) Score() int {
	return a.score
}

// Modifier returns the score modifier for the current score value
func (s *Ability) Modifier() int {
	score := float64(s.Score())
	value := (score - 10) / 2
	return int(math.Floor(value))
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
