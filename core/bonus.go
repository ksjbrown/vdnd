package core

// Bonus represents some effect that provides a scalar bonus to some property
// A key should be defined, so that the bonus can be uniquely identified, and not applied more than once
type Bonus interface {
	GetBonusKey() any
	GetBonusValue() int
}

// SupportsBonuses defines a sensible set of operations for supporting bonuses
type SupportsBonuses interface {
	AddBonus(Bonus)
	RemoveBonus(Bonus)
	HasBonus(Bonus) bool
	GetBonus(key any) Bonus
}

// ValueBonusKey provides a simple embeddable implementation for GetBonusKey()
// It is not really intented to be used on its own.
type ValueBonusKey struct {
	Key any
}

func NewValueBonusKey(key any) *ValueBonusKey {
	return &ValueBonusKey{Key: key}
}

func (p *ValueBonusKey) GetBonusKey() any {
	return p.Key
}

// ValueBonus implements the Bonus interface, and returns the stored Value via GetBonusValue()
type ValueBonus struct {
	ValueBonusKey
	Value int
}

func NewValueBonus(key any, value int) *ValueBonus {
	return &ValueBonus{
		ValueBonusKey: ValueBonusKey{Key: key},
		Value:         value,
	}
}

func (b *ValueBonus) GetBonusValue() int {
	return b.Value
}

// FuncBonus implements the Bonus interface, and evaluates ValueFunc when calling GetBonusValue()
type FuncBonus struct {
	ValueBonusKey
	Func func() int
}

func NewFuncBonus(key any, valueFunc func() int) *FuncBonus {
	return &FuncBonus{
		ValueBonusKey: ValueBonusKey{Key: key},
		Func:          valueFunc,
	}
}

func (b *FuncBonus) GetBonusValue() int {
	return b.Func()
}

// BonusMap is a basic, map backed store of Bonuses.
//
// Bonuses are uniquely identified by GetBonusKey()
type BonusMap struct {
	bonuses map[any]Bonus
}

func NewBonusMap() *BonusMap {
	return &BonusMap{
		bonuses: make(map[any]Bonus),
	}
}

func (s *BonusMap) AddBonus(b Bonus) {
	key := b.GetBonusKey()
	s.bonuses[key] = b
}

func (s *BonusMap) RemoveBonus(b Bonus) {
	key := b.GetBonusKey()
	delete(s.bonuses, key)
}

func (s *BonusMap) HasBonus(b Bonus) bool {
	key := b.GetBonusKey()
	_, found := s.bonuses[key]
	return found
}

func (s *BonusMap) GetBonus(key any) Bonus {
	return s.bonuses[key]
}
