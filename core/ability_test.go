package core

import (
	"testing"
)

func TestAbilitiesModifier(t *testing.T) {
	tcs := []struct {
		name     string
		score    int
		modifier int
	}{
		{score: 7, modifier: -2},
		{score: 8, modifier: -1},
		{score: 9, modifier: -1},
		{score: 10, modifier: 0},
		{score: 11, modifier: 0},
		{score: 12, modifier: 1},
		{score: 13, modifier: 1},
		{score: 14, modifier: 2},
	}
	kind := AbilityKind(0)
	for _, tc := range tcs {
		ability := NewAbility(kind, tc.score)
		modifier := ability.Modifier()
		if tc.modifier != modifier {
			t.Errorf("ability score '%v' expects modifier '%v', got '%v'", tc.score, tc.modifier, modifier)
		}
	}
}
