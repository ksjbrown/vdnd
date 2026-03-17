package core

type Class struct {
	Type ClassType
	Levels []ClassLevel
	// Progression
}

// ClassType represents the general classification of a Character.
type ClassType int

const (
	Barbarian ClassType = 0x1
	Bard ClassType = 0x2
	Cleric ClassType = 0x3
	Druid ClassType = 0x4
	Fighter ClassType = 0x5
	Monk ClassType = 0x6
	Paladin ClassType = 0x7
	Ranger ClassType = 0x8
	Rogue ClassType = 0x9
	Sorcerer ClassType = 0xA
	Warlock ClassType = 0xB
	Wizard ClassType = 0xC
)

type ClassLevel struct {

}
