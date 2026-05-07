package core

import "testing"

func Test_CharacterAbility(t *testing.T) {
	c := &Character{
		BaseAbilities: Abilities{
			Strength: 17,
		},
	}
	mutator := &DeltaMutatorFunc{func() int {
		return 1
	}}
	c.AbilitiesMutator.SetDeltaMutator(
		AbilityKindStrength,
		mutator,
	)
	ability := c.GetAbility(AbilityKindStrength)
	if ability != Ability(18) {
		t.Errorf("expected 18, got %v", ability)
	}

}
