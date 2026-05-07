package core

type Ability struct {
	baseValue int
	mutators  []AbilityMutator
}

func NewAbility(value int) *Ability {
	// TODO: bounds check
	return &Ability{
		baseValue: value,
		mutators:  make([]AbilityMutator, 0),
	}
}

func (a *Ability) AddMutator(mutator AbilityMutator) {
	a.mutators = append(a.mutators, mutator)
}

func (a *Ability) Value() int {
	delta := 0
	for _, mutator := range a.mutators {
		delta += mutator.GetAbilityValueDelta()
	}
	return a.baseValue + delta
}


func (a *Ability) ModifierBonus() int {
	// offset from 10, divided by 2, rounded down
    value := a.Value() - 10
    if value >= 0 {
        return value / 2
    }
    return (value - 1) / 2
}

