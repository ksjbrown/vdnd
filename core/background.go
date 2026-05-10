package core

type SupportsBackground interface {
	GetBackground() BackgroundKind
	GetBackgroundFeat() Feat
	GetBackgroundSkillProficiencies() []SkillKind
	GetBackgroundToolProficiencies() []ToolKind
}

type BackgroundData struct {
	Kind BackgroundKind
	Feat FeatData
}

// Background implements the SupportsBackground interface.
//
// It defines the background of the character,
// and the selected feats, ability boosts, etc.
type Background struct {
	kind BackgroundKind
	feat Feat
	skillProficiencies []SkillKind
	toolProficiencies []ToolKind
}

func (b *Background) GetBackground() BackgroundKind {
	return b.kind
}

func (b *Background) GetBackgroundFeat() Feat {
	return b.feat
}

type BackgroundKind int

const (
	BackgroundKindAcolyte     BackgroundKind = 0x01
	BackgroundKindArtisan     BackgroundKind = 0x02
	BackgroundKindCharlatan   BackgroundKind = 0x03
	BackgroundKindCriminal    BackgroundKind = 0x04
	BackgroundKindEntertainer BackgroundKind = 0x05
	BackgroundKindFarmer      BackgroundKind = 0x06
	BackgroundKindGuard       BackgroundKind = 0x07
	BackgroundKindGuide       BackgroundKind = 0x08
	BackgroundKindHermit      BackgroundKind = 0x09
	BackgroundKindMerchant    BackgroundKind = 0x0A
	BackgroundKindNoble       BackgroundKind = 0x0B
	BackgroundKindSage        BackgroundKind = 0x0C
	BackgroundKindSailor      BackgroundKind = 0x0D
	BackgroundKindScribe      BackgroundKind = 0x0E
	BackgroundKindSoldier     BackgroundKind = 0x0F
	BackgroundKindWayfarer    BackgroundKind = 0x10
)
