package core

type Ability int

func (a Ability) ModifierBonus() int {
	// offset from 10, divided by 2, rounded down
	value := int(a) - 10
	if value >= 0 {
		return value / 2
	}
	return (value - 1) / 2
}
