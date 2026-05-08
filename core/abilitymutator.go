package core

type AbilityMutatorUndo = func()

type AbilityMutator interface {
	GetAbilityValueBonus() int
}

type FixedBonusAbilityMutator struct {
	delta int
}

func (m *FixedBonusAbilityMutator) GetAbilityValueBonus() int {
	return m.delta
}
