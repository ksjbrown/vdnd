package core

type CharacterMutators struct {
	AbilitiesMutator
	AttackRollMutator
	ArmorClassMutator
}

type AbilitiesMutator struct {
	DeltaMutatorMap[AbilityKind]
}

type ArmorClassMutator struct {
	DeltaMutatorComposite
}

type AttackRollMutator struct {
	DeltaMutatorComposite
}
