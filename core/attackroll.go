package core

type SupportsAttackRoll interface {
	GetAttackRoll() int
}

type AttackRoll struct {
	SupportsBonuses
}


