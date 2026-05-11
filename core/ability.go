package core

type SupportsEffectsAbilities interface {
	AddAbilitiesEffect(Effect[SupportsAbilities])
	RemoveAbilitiesEffect(AbilityKind, Effect[SupportsAbility])
}

type SupportsAbilities interface {
	SupportsEffectsAbilities
	GetAbility(AbilityKind) SupportsAbility
	GetStrength() SupportsAbility
	GetDexterity() SupportsAbility
	GetConstitution() SupportsAbility
	GetIntelligence() SupportsAbility
	GetWisdom() SupportsAbility
	GetCharisma() SupportsAbility
}

type AbilitiesData struct {
	Strength     int
	Dexterity    int
	Constitution int
	Wisdom       int
	Intelligence int
	Charisma     int
}

// Abilities holds the six ability scores for a character.
type Abilities struct {
	abilities map[AbilityKind]*Ability
}

func NewAbilities(data *AbilitiesData) *Abilities {
	a := make(map[AbilityKind]*Ability)
	a[AbilityKindStrength] = NewAbility(AbilityKindStrength, data.Strength)
	a[AbilityKindDexterity] = NewAbility(AbilityKindDexterity, data.Dexterity)
	a[AbilityKindConstitution] = NewAbility(AbilityKindConstitution, data.Constitution)
	a[AbilityKindWisdom] = NewAbility(AbilityKindWisdom, data.Wisdom)
	a[AbilityKindIntelligence] = NewAbility(AbilityKindIntelligence, data.Intelligence)
	a[AbilityKindCharisma] = NewAbility(AbilityKindCharisma, data.Charisma)
	return &Abilities{
		abilities: a,
	}
}

func (a *Abilities) GetAbility(kind AbilityKind) *Ability {
	return a.abilities[kind]
}

func (a *Abilities) GetStrength() *Ability {
	return a.GetAbility(AbilityKindStrength)
}

func (a *Abilities) GetDexterity() *Ability {
	return a.GetAbility(AbilityKindDexterity)
}

func (a *Abilities) GetConstitution() *Ability {
	return a.GetAbility(AbilityKindConstitution)
}

func (a *Abilities) GetWisdom() *Ability {
	return a.GetAbility(AbilityKindWisdom)
}

func (a *Abilities) GetIntelligence() *Ability {
	return a.GetAbility(AbilityKindIntelligence)
}

func (a *Abilities) GetCharisma() *Ability {
	return a.GetAbility(AbilityKindCharisma)
}

type SupportsEffectsAbility interface {
	AddAbilityEffect(Effect[SupportsAbility])
	RemoveAbilityEffect(Effect[SupportsAbility])
}

type SupportsAbility interface {
	GetAbilityKind() AbilityKind
	GetAbilityBaseScore() int
	GetAbilityScore() int
	GetAbilityModifier() int
}

// Ability represents a single Ability score (STR, DEX, etc.)
//
// An Ability Score consists of a BaseScore, as well as any number of Bonuses, like:
//
// - Character Creation Ability Point distributions
// - Feats
// - Temporary Buffs or Debuffs
type Ability struct {
	SupportsBonuses
	kind      AbilityKind
	baseScore int
}

func NewAbility(kind AbilityKind, baseScore int) *Ability {
	return &Ability{
		SupportsBonuses: NewBonusMap(),
		kind:            kind,
		baseScore:       baseScore,
	}
}

func (a *Ability) GetAbilityKind() AbilityKind {
	return a.kind
}

func (a *Ability) GetAbilityBaseScore() int {
	return a.baseScore
}

func (a *Ability) GetAbilityScore() int {
	return a.GetAbilityBaseScore() + a.GetCombinedBonusValue()
}

func (a *Ability) GetAbilityModifier() int {
	// score offset from 10, divided by 2, rounded down
	offset := a.GetAbilityScore() - 10
	if offset >= 0 {
		return offset / 2
	}
	return (offset - 1) / 2
}

type AbilityKind int

const (
	AbilityKindStrength     AbilityKind = 0x1
	AbilityKindDexterity    AbilityKind = 0x2
	AbilityKindConstitution AbilityKind = 0x3
	AbilityKindWisdom       AbilityKind = 0x4
	AbilityKindIntelligence AbilityKind = 0x5
	AbilityKindCharisma     AbilityKind = 0x6
)
