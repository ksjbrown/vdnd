package core

type AbilityMutatorUndo = func()

type AbilityMutator interface {
	GetAbilityValueDelta() int
}

type FixedDeltaAbilityMutator struct {
	delta int
}

func (m *FixedDeltaAbilityMutator) GetAbilityValueDelta() int {
	return m.delta
}
