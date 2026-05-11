package core

import (
	"testing"
)

func TestAbilitiesModifier(t *testing.T) {
	tcs := []struct {
		name     string
		value    int
		modifier int
	}{
		{value: 7, modifier: -2},
		{value: 8, modifier: -1},
		{value: 9, modifier: -1},
		{value: 10, modifier: 0},
		{value: 11, modifier: 0},
		{value: 12, modifier: 1},
		{value: 13, modifier: 1},
		{value: 14, modifier: 2},
	}
	for _, tc := range tcs {
		ability := NewAbility(0, tc.value)
		modifier := ability.GetAbilityModifier()
		if tc.modifier != modifier {
			t.Errorf("ability score '%v' expects modifier '%v', got '%v'", tc.value, tc.modifier, modifier)
		}
	}
}
