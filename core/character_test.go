package core

import "testing"

func Test_CharacterAbility(t *testing.T) {
	c := &Character{
		BaseAbilities: Abilities{
			Strength: 17,
		},
	}
	dp := &BonusProviderFunc{func() int {
		return 1
	}}
	c.modifiers.Abilities.AppendBonusProvider(
		AbilityKindStrength,
		dp,
	)
	ability := c.GetAbility(AbilityKindStrength)
	if ability != Ability(18) {
		t.Errorf("expected 18, got %v", ability)
	}
	c.modifiers.Abilities.AppendBonusProvider(
		AbilityKindStrength,
		dp,
	)
	ability = c.GetAbility(AbilityKindStrength)
	if ability != Ability(19) {
		t.Errorf("expected 19, got %v", ability)
	}
}
