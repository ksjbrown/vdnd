package core

type Feat struct {
	Type FeatType
	data any // FeatType specific payload that may exist, e.g. AbilityScoreImprovement must choose an ability
}

type FeatType int

const (
	AbilityScoreImprovement   FeatType = 0x01
	Actor                     FeatType = 0x07
	Alert                     FeatType = 0x08
	Archery                   FeatType = 0x09
	Athlete                   FeatType = 0x0A
	BlindFighting             FeatType = 0x0B
	BoonOfCombatProwess       FeatType = 0x0B
	BoonOfDimensionalTravel   FeatType = 0x0B
	BoonOfEnergyResistance    FeatType = 0x0B
	BoonOfFate                FeatType = 0x0B
	BoonOfFortitude           FeatType = 0x0B
	BoonOfIrresistableOffence FeatType = 0x0B
	BoonOfRecovery            FeatType = 0x0B
	BoonOfSkill               FeatType = 0x0B
	BoonOfSpeed               FeatType = 0x0B
	BoonOfSpellRecall         FeatType = 0x0B
	BoonOfTheNightSpirit      FeatType = 0x0B
	BoonOfTruesight           FeatType = 0x0B
	Charger                   FeatType = 0x0B
	Chef                      FeatType = 0x0B
	Crafter                   FeatType = 0x0B
	CrossbowExpert            FeatType = 0x0B
	Crusher                   FeatType = 0x0B
	Defence                   FeatType = 0x0B
	DefenciveDualist          FeatType = 0x0B
	DualWielder               FeatType = 0x0B
	Dueling                   FeatType = 0x0B
	Durable                   FeatType = 0x0B
	ElementalAdept            FeatType = 0x0B
	FeyTouched                FeatType = 0x0B
	Grappler                  FeatType = 0x0B
	GreatWeaponFighting       FeatType = 0x0B
	GreatWeaponmaster         FeatType = 0x0B
	Healer                    FeatType = 0x0B
	HeavilyArmored            FeatType = 0x0B
	HeavyArmorMaster          FeatType = 0x0B
	InspiringLeader           FeatType = 0x0B
	Interception              FeatType = 0x0B
	KeenMind                  FeatType = 0x0B
	LightlyArmored            FeatType = 0x0B
	Lucky                     FeatType = 0x0B
	MageSlayer                FeatType = 0x0B
	MagicInitiate             FeatType = 0x0B
	MartialWeaponTraining     FeatType = 0x0B
	MediumArmorMaster         FeatType = 0x0B
	ModeratelyArmored         FeatType = 0x0B
	MountedCombatant          FeatType = 0x0B
	Musician                  FeatType = 0x0B
	Observant                 FeatType = 0x0B
	Piercer                   FeatType = 0x0B
	Poisoner                  FeatType = 0x0B
	PolearmMaster             FeatType = 0x0B
	Protection                FeatType = 0x0B
	Resilient                 FeatType = 0x0B
	RitualCaster              FeatType = 0x0B
	SavageAttacker            FeatType = 0x0B
	Sentinel                  FeatType = 0x0B
	ShadowTouched             FeatType = 0x0B
	Sharpshooter              FeatType = 0x0B
	ShieldMaster              FeatType = 0x0B
	Skilled                   FeatType = 0x0B
	SkillExpert               FeatType = 0x0B
	Skulker                   FeatType = 0x0B
	Slasher                   FeatType = 0x0B
	Speedy                    FeatType = 0x0B
	SpellSniper               FeatType = 0x0B
	TavernBrawler             FeatType = 0x00
	Telekinetic               FeatType = 0x00
	Telepathic                FeatType = 0x00
	ThrownWeaponFighting      FeatType = 0x00
	Tough                     FeatType = 0x00
	TwoWeaponFighting         FeatType = 0x00
	UnarmedFighting           FeatType = 0x00
	WarCaster                 FeatType = 0x00
	WeaponMaster              FeatType = 0x00
)
