package core

type Effect interface {
	ApplyEffectTo(target any)
}

type TargetedEffect[T any] interface {
	ApplyEffectTo(target T)
}

