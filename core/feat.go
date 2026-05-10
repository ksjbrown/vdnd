package core

import "encoding/json"

type SupportsFeats interface {
	AddFeat(Feat)

}

type Feat interface {
	GetFeatKind() FeatKind
	IsRepeatable() bool
	GetEffects() []Effect
}


type FeatData struct {
	Kind FeatKind
	Data json.RawMessage
}

func NewFeat(data *FeatData) Feat {
	panic("unimplememnted")
}

// FeatMap implements SupportsFeats
//
// It supports storing feats, and handles feats that are repeatable
type FeatMap struct {

}

// -- Implementations --

type AbilityScoreImprovementFeat struct {
	Kind AbilityKind
	Value int
}

func (f *AbilityScoreImprovementFeat) GetFeatKind() FeatKind {
	return FeatKindAbilityScoreImprovement
}



func (f *AbilityScoreImprovementFeat) GetBonus() Bonus {
	bonus := NewValueBonus(
		FeatKindAbilityScoreImprovement,
		f.Value,
	)
	return bonus

}
