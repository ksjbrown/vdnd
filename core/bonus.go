package core

// Bonus represents some effect that provides a scalar bonus to some property
// A key should be defined, so that the bonus can be uniquely identified, and not applied more than once
type Bonus interface {
	GetBonusKey() any
	GetBonusValue() int
}

// BonusKeyProvider provides a simple embeddable implementation for GetBonusKey()
// It is not really intented to be used on its own.
type BonusKeyProvider struct {
	Key any
}

func NewBonusKeyProvider(key any) *BonusKeyProvider {
	return &BonusKeyProvider{Key: key}
}

func (p *BonusKeyProvider) GetBonusKey() any {
	return p.Key
}

// DeltaBonus implements the Bonus interface, and returns the stored Value via GetBonusValue()
type SimpleBonus struct {
	BonusKeyProvider
	Value int
}

func NewSimpleBonus(key any, value int) *SimpleBonus {
	return &SimpleBonus{
		BonusKeyProvider: BonusKeyProvider{Key: key},
		Value:            value,
	}
}

func (b *SimpleBonus) GetBonusValue() int {
	return b.Value
}

// FuncBonus implements the Bonus interface, and evaluates ValueFunc when calling GetBonusValue()
type FuncBonus struct {
	BonusKeyProvider
	ValueFunc func() int
}

func NewFuncBonus(key any, valueFunc func() int) *FuncBonus {
	return &FuncBonus{
		BonusKeyProvider: BonusKeyProvider{Key: key},
		ValueFunc:        valueFunc,
	}
}

func (b *FuncBonus) GetBonusValue() int {
	return b.ValueFunc()
}

// SupportsBonuses defines a sensible set of operations for supporting bonuses
type SupportsBonuses interface {
	AddBonus(Bonus)
	GetBonus(key any) Bonus
	HasBonus(Bonus) bool
	RemoveBonus(Bonus)
	ComputeBonuses() int
}

// SimpleBonusSupporter is a basic, map backed store of Bonuses.
//
// Bonuses are uniquely identified by GetBonusKey()
type SimpleBonusSupporter struct {
	bonuses map[any]Bonus
}

func NewSimpleBonusSupporter() *SimpleBonusSupporter {
	return &SimpleBonusSupporter{
		bonuses: make(map[any]Bonus),
	}
}

func (s *SimpleBonusSupporter) AddBonus(b Bonus) {
	key := b.GetBonusKey()
	s.bonuses[key] = b
}

func (s *SimpleBonusSupporter) GetBonus(key any) Bonus {
	return s.bonuses[key]
}

func (s *SimpleBonusSupporter) HasBonus(b Bonus) bool {
	key := b.GetBonusKey()
	_, found := s.bonuses[key]
	return found
}

func (s *SimpleBonusSupporter) RemoveBonus(b Bonus) {
	key := b.GetBonusKey()
	delete(s.bonuses, key)
}

func (s *SimpleBonusSupporter) ComputeBonuses() int {
	bonus := 0
	for _, b := range s.bonuses {
		bonus += b.GetBonusValue()
	}
	return bonus
}

// CompositeBonus uses the SimpleBonusSupporter to compose multiple bonuses under the Bonus interface.
type CompositeBonus struct {
	BonusKeyProvider
	SimpleBonusSupporter
}

func NewCompositeBonus(key any) *CompositeBonus {
	return &CompositeBonus{
		BonusKeyProvider:     *NewBonusKeyProvider(key),
		SimpleBonusSupporter: *NewSimpleBonusSupporter(),
	}
}

func (b *CompositeBonus) GetBonusValue() int {
	return b.ComputeBonuses()
}
