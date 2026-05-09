package core

import "encoding/json"

type Feat interface {
	GetFeatKind() FeatKind
	IsRepeatable() bool
}

type SupportsFeats interface {
	AddFeat(Feat)
}

type FeatData struct {
	Kind FeatKind
	Data json.RawMessage
}

func NewFeat(data *FeatData) Feat {
	panic("unimplememnted")
}

// -- Implementations --

// BaseFeat provides the minimal required data for a Feat,
// as well as no-op implementations for every possible place a feat can manipulate game state.
type BaseFeat struct {
	kind       FeatKind
	repeatable bool
}

func NewBaseFeat(kind FeatKind, repeatable bool) *BaseFeat {
	return &BaseFeat{
		kind: kind,
		repeatable: repeatable,
	}
}

func (f *BaseFeat) GetFeatKind() FeatKind {
	return f.kind
}

func (f *BaseFeat) IsRepeatable() bool {
	return f.repeatable
}

type AbilityScoreImprovementFeat struct {
	Feat
}

func NewAbilityScoreImprovementFeat() *AbilityScoreImprovementFeat {
	return &AbilityScoreImprovementFeat{
		Feat: &BaseFeat{
			kind:       FeatKindAbilityScoreImprovement,
			repeatable: true,
		},
	}
}

func (f *AbilityScoreImprovementFeat) GetBonus() Bonus {
	bonus := NewSimpleBonus(FeatKindAbilityScoreImprovement, f.value)

}
