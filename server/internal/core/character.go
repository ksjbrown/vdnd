package core

// Character represents any person, creature, or actor in the game, that can perform actions, be interacted with etc.
//
// Should contain all data required for character functionality.
//
// Some base types are extended in Character context, e.g. Feat -> CharacterFeat.
// This is usually done because some selection or choice is provided by the base type.
// The instance user in the Character will reflect the choice the Character owner made.
type Character struct {
	Name       string
	Level      CharacterLevel
	Abilities
	Background
	Class      Class
	Origin     Origin
	HP         Health
	AC         ArmorClass
}

// CharacterBackground defines the background of the character,
// and the selected feats, ability boosts, etc.
type CharacterBackground struct {
	Name         BackgroundType
	AbilityBonus map[AbilityType]AbilityScore
	Feat         FeatType
}

type CharacterLevel int
