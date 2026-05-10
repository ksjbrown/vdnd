package core

type SupportsAlignment interface {
	GetAlignment() AlignmentKind
}

type Alignment struct {
	Kind AlignmentKind
}

func NewAlignment(kind AlignmentKind) *Alignment {
	return &Alignment{
		Kind: kind,
	}
}

func (a *Alignment) GetAlignment() AlignmentKind {
	return a.Kind
}

type AlignmentKind int

const (
	AlignmentLawfulGood     AlignmentKind = 0x1
	AlignmentNeutralGood    AlignmentKind = 0x2
	AlignmentChaoticGood    AlignmentKind = 0x3
	AlignmentLawfulNeutral  AlignmentKind = 0x4
	AlignmentTrueNeutral    AlignmentKind = 0x5
	AlignmentChaoticNeutral AlignmentKind = 0x6
	AlignmentLawfulEvil     AlignmentKind = 0x7
	AlignmentNeutralEvil    AlignmentKind = 0x8
	AlignmentChaoticEvil    AlignmentKind = 0x9
)
