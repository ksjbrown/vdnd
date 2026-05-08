package core

type BonusProviderMap[T comparable] interface {
	GetBonusProvider(key T) BonusProvider
	GetBonus(key T) int
}

type CompositeBonusProviderMap[T comparable] struct {
	dps map[T]*CompositeBonusProvider
}

func (m *CompositeBonusProviderMap[T]) GetBonusProvider(key T) BonusProvider {
	return m.dps[key]
}

func (m *CompositeBonusProviderMap[T]) GetBonus(key T) int {
	mutator := m.GetBonusProvider(key)
	if mutator == nil {
		return 0
	}
	return mutator.GetBonus()
}

func (m *CompositeBonusProviderMap[T]) AppendBonusProvider(key T, dp BonusProvider) {
	if m.dps == nil {
		m.dps = make(map[T]*CompositeBonusProvider)
	}
	cdp := m.dps[key]
	if cdp == nil {
		cdp = &CompositeBonusProvider{}
		m.dps[key] = cdp
	}
	cdp.AppendBonusProvider(dp)
}

