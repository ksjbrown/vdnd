package core

type Class struct {
	kind ClassKind
}

func (c *Class) Kind() ClassKind {
	return c.kind
}

// ClassKind represents the general classification of a Character.
type ClassKind int

const (
	ClassKindBarbarian ClassKind = 0x1
	ClassKindBard      ClassKind = 0x2
	ClassKindCleric    ClassKind = 0x3
	ClassKindDruid     ClassKind = 0x4
	ClassKindFighter   ClassKind = 0x5
	ClassKindMonk      ClassKind = 0x6
	ClassKindPaladin   ClassKind = 0x7
	ClassKindRanger    ClassKind = 0x8
	ClassKindRogue     ClassKind = 0x9
	ClassKindSorcerer  ClassKind = 0xA
	ClassKindWarlock   ClassKind = 0xB
	ClassKindWizard    ClassKind = 0xC
)
