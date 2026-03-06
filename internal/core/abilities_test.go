package core

import (
	"testing"
)

func Test_Abilities_Modifier(t *testing.T) {
	tcs := []struct{
		Name string
		Score int
		Modifier int
	} {
		{ Score: 10, Modifier: 0 },
		{ Score: 11, Modifier: 0 },
		{ Score: 12, Modifier: 1 },
		{ Score: 9, Modifier: -1 },
		{ Score: 8, Modifier: -1 },
		{ Score: 7, Modifier: -2 },
	}
	for _, tc := range tcs {
		score := AbilityScore(tc.Score)
		modifier := score.Modifier()
		if tc.Modifier != modifier {
			t.Errorf("unexpected")
		}
	}
}
