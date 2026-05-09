package core

// CharacterModifiers is a collection of all mutators that are associated with a Character
type CharacterModifiers struct {
	Abilities   AbilityModifiers
	AttackRolls AttackRollModifiers
	ArmorClass  ArmorClassModifiers
}

type AbilityModifiers struct {
	CompositeBonusProviderMap[AbilityKind]
}

type ArmorClassModifiers struct {
	CompositeBonusProvider
}

type AttackRollModifiers struct {
	CompositeBonusProvider
	CompositeAdvantageProvider
}
