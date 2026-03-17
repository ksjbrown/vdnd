package core

// Background defines the background of the character,
// and the selected feats, ability boosts, etc.
type Background struct {
	Type          BackgroundType
	Feat          Feat
	Proficiencies BackgroundProficiencies
}

type BackgroundType int

const (
	Acolyte     BackgroundType = 0x01
	Artisan     BackgroundType = 0x02
	Charlatan   BackgroundType = 0x03
	Criminal    BackgroundType = 0x04
	Entertainer BackgroundType = 0x05
	Farmer      BackgroundType = 0x06
	Guard       BackgroundType = 0x07
	Guide       BackgroundType = 0x08
	Hermit      BackgroundType = 0x09
	Merchant    BackgroundType = 0x0A
	Noble       BackgroundType = 0x0B
	Sage        BackgroundType = 0x0C
	Sailor      BackgroundType = 0x0D
	Scribe      BackgroundType = 0x0E
	Soldier     BackgroundType = 0x0F
	Wayfarer    BackgroundType = 0x10
)

type BackgroundProficiencies struct {
	Skills [BackgroundProficiencySkillCount]SkillType
	Tools  [BackgroundProficiencyToolCount]ToolType
}

const BackgroundProficiencySkillCount = 2
const BackgroundProficiencyToolCount = 1
