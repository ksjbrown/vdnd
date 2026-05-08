package core

import "encoding/json"

type FeatData struct {
	Kind FeatKind
	Data json.RawMessage
}

type Feat interface {
}

func NewFeat(data *FeatData) Feat {
	panic("unimplememnted")
}

// -- Implementations --

type AbilityScoreImprovementFeat struct {
	abilityKind AbilityKind
	value int
}

func (f *AbilityScoreImprovementFeat) MutateAbility(a *AbilityModifiers) {
	
}
