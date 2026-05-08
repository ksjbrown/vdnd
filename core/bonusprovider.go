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
