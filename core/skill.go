package core

// SkillKind represents a skill that a Character might have.
//
// It doesn't have a value, a Character either has this skill proficiency, or he does not.
type SkillKind int

const (
	SkillKindAcrobatics     SkillKind = 0x01
	SkillKindAnimalHandling SkillKind = 0x02
	SkillKindArcana         SkillKind = 0x03
	SkillKindAthletics      SkillKind = 0x04
	SkillKindDeception      SkillKind = 0x05
	SkillKindHistory        SkillKind = 0x06
	SkillKindInsight        SkillKind = 0x07
	SkillKindIntimidation   SkillKind = 0x08
	SkillKindInvestigation  SkillKind = 0x09
	SkillKindMedicine       SkillKind = 0x0A
	SkillKindNature         SkillKind = 0x0B
	SkillKindPerception     SkillKind = 0x0C
	SkillKindPerformance    SkillKind = 0x0D
	SkillKindPersuation     SkillKind = 0x0E
	SkillKindReligion       SkillKind = 0x0F
	SkillKindSleightOfHand  SkillKind = 0x10
	SkillKindStealth        SkillKind = 0x11
	SkillKindSurvival       SkillKind = 0x12
)
