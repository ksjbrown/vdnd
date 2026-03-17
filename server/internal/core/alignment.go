package core

type Alignment int

const (
	LawfulGood     Alignment = 0x1
	NeutralGood    Alignment = 0x2
	ChaoticGood    Alignment = 0x3
	LawfulNeutral  Alignment = 0x4
	TrueNeutral    Alignment = 0x5
	ChaoticNeutral Alignment = 0x6
	LawfulEvil     Alignment = 0x7
	NeutralEvil    Alignment = 0x8
	ChaoticEvil    Alignment = 0x9
)
