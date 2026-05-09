package core

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
		SupportsBonuses: NewSimpleBonusSupporter(),
		BaseScore:       baseScore,
	}
}

func (a *Ability) Score() int {
	return a.BaseScore + a.ComputeBonuses()
}

func (a *Ability) Modifier() int {
	// offset from 10, divided by 2, rounded down
	score := a.Score()
	offset := int(score) - 10
	if offset >= 0 {
		return offset / 2
	}
	return (offset - 1) / 2
}
