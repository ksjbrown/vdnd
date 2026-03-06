package core

import "math"

type AbilityScore int

func (s AbilityScore) Modifier() int {
	return int(math.Floor(float64((s - 10) / 2)))
}

type AbilityScores struct {
	Str AbilityScore
	// TODO: complete
}
