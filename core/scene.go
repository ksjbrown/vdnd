package core

import (
	"fmt"
)

// Scene defines a generic contract for processing player actions,
// and applying results to those actions.
//
// Primary mechanism for progressing a scene is thru Actions.
// An Action is provided to the scene via the Accept method.
// The scene may perform validation on the action if required.
// Otherwise, the action

type Scene interface {
	GetCharacters() []*Character
	GetCharacter(CharacterID) *Character
	Accept(Action) error
}

type CombatScene struct {
}

func ProcessAttack(scene Scene, attackerID, targetID CharacterID, attack Action) error {

	attacker := scene.GetCharacter(attackerID)
	if attacker == nil {
		return fmt.Errorf("no attacker with id: %v", attackerID)
	}

	target := scene.GetCharacter(targetID)
	if target == nil {
		return fmt.Errorf("no target with id: %v", targetID)
	}

	//
}
