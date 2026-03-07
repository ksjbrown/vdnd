package core

// Actor represents some entity in the world that can Act.
//
// This may be an NPC, player controlled character, etc..
type Actor struct {
	*Abilities
}
