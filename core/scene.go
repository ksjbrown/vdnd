package core

// Scene defines a generic contract for processing player actions,
// and applying results to those actions.
//
// Primary mechanism for progressing a scene is thru Actions.
// An Action is provided to the scene via the Accept method.
// The scene may perform validation on the action if required.
// Otherwise, the action 

type Scene interface {
	GetEntities() []*Entity
	GetEntity(EntityID) *Entity
	GetCharacters() []*Character
	GetCharacter(CharacterID) *Character
	Accept(Action) error
}

type CombatScene struct {

}

// Potential Scenes:
//
// - CombatScene
// - DialogueScene

