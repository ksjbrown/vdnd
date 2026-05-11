package core

import "iter"

// Bonus represents some effect that provides a scalar bonus to some property
// A key should be defined, so that the bonus can be uniquely identified, and not applied more than once
type Bonus interface {
	GetBonusKey() any
	GetBonusValue() int
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

// SupportsBonuses defines a sensible set of operations for supporting bonuses.
//
// SupportsBonuses can add bonuses for the same key, but only the best value for a given key is included in the combined bonus method.
// Adding the exact same bonus is a no-op
type SupportsBonuses interface {
	AddBonus(Bonus)
	RemoveBonus(Bonus)
	GetCombinedBonusValue() int
	IterateBonuses() iter.Seq[Bonus]
}

// BonusMap is a basic, map backed store of Bonuses.
//
// Bonuses are uniquely identified by GetBonusKey()
type BonusMap struct {
	bonuses map[Bonus]struct{}
}

func NewBonusMap() *BonusMap {
	return &BonusMap{
		bonuses: make(map[Bonus]struct{}),
	}
}

func (s *BonusMap) AddBonus(b Bonus) {
	s.bonuses[b] = struct{}{}
}

func (s *BonusMap) RemoveBonus(b Bonus) {
	delete(s.bonuses, b)
}

func (s *BonusMap) GetCombinedBonusValue() int {
	value := 0
	for bonus := range s.bonuses {
		value += bonus.GetBonusValue()
	}
	return value
}

func (s *BonusMap) IterateBonuses() iter.Seq[Bonus] {
	return func(yield func(Bonus) bool) {
		for bonus := range s.bonuses {
			if !yield(bonus) {
				return
			}
		}
	}
}

type BonusKind int

const (
	BonusKindArmorClassDexterity BonusKind = 0x01
)
