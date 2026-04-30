package core

type Alignment int

const (
	AlignmentLawfulGood     Alignment = 0x1
	AlignmentNeutralGood    Alignment = 0x2
	AlignmentChaoticGood    Alignment = 0x3
	AlignmentLawfulNeutral  Alignment = 0x4
	AlignmentTrueNeutral    Alignment = 0x5
	AlignmentChaoticNeutral Alignment = 0x6
	AlignmentLawfulEvil     Alignment = 0x7
	AlignmentNeutralEvil    Alignment = 0x8
	AlignmentChaoticEvil    Alignment = 0x9
)
