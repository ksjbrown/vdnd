package core

// BonusProvider represents any type that provides a method ModifierBonus()
type BonusProvider interface {
	GetBonus() int
}

type CompositeBonusProvider struct {
	dps []BonusProvider
}

func (m *CompositeBonusProvider) GetBonus() int {
	delta := 0
	for _, dp := range m.dps {
		delta += dp.GetBonus()
	}
	return delta
}

func (m *CompositeBonusProvider) AppendBonusProvider(dp BonusProvider) {
	m.dps = append(m.dps, dp)
}

type BonusProviderFunc struct {
	fn func() int
}

func (m *BonusProviderFunc) GetBonus() int {
	return m.fn()
}

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
