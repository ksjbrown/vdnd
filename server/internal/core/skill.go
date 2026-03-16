package core

type Skill struct {
	Type SkillType
}

// SkillType represents a skill that a Character might have.
//
// It doesn't have a value, a Character either has this skill proficiency, or he does not.
type SkillType int

const (
	Acrobatics     SkillType = 0x01
	AnimalHandling SkillType = 0x02
	Arcana         SkillType = 0x03
	Athletics      SkillType = 0x04
	Deception      SkillType = 0x05
	History        SkillType = 0x06
	Insight        SkillType = 0x07
	Intimidation   SkillType = 0x08
	Investigation  SkillType = 0x09
	Medicine       SkillType = 0x0A
	Nature         SkillType = 0x0B
	Perception     SkillType = 0x0C
	Performance    SkillType = 0x0D
	Persuation     SkillType = 0x0E
	Religion       SkillType = 0x0F
	SleightOfHand  SkillType = 0x10
	Stealth        SkillType = 0x11
	Survival       SkillType = 0x12
)
