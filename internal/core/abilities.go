package core

import "math"

type Ability int

func (s Ability) Score() int {
	return int(math.Floor(float64(s - 10) / 2))
}

type Abilities struct {
	Str Ability
	// TODO: complete
}
