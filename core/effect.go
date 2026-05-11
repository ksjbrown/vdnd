package core

// TODO: some open questions about architecture:
//
// - Should Effect handle all effects, including permanent effects? like feat stat improvements?
// - Relationship to other "effects"?
//   - Bonus, Multiplier/Factor
//     - "Raw" numerical modifications
//   - Effect
//     - A package of bonuses.
//     - Enompasses the overall result or the effect, as well logic to add, remove etc.
//   - EffectProviders (Feats, Items, etc)
//     - Can provide any number of targeted effects.

// Effect defines the modifications to the target T that occur because of this effect.
//
// The effect is responsible for manipulating the state of T via T's API.
type Effect[T any] interface {
	GetEffectKey() any
	ActivateEffect(target T)
	DeactivateEffect(target T)
	IsExpired() bool
}

type FuncEffect[T any] struct {
	key        func() any
	activate   func(target T)
	deactivate func(target T)
	expired    func() bool
}

func NewFuncEffect[T any](key func() any, activate func(T), deactivate func(T), expired func() bool) *FuncEffect[T] {
	return &FuncEffect[T]{
		activate:   activate,
		deactivate: deactivate,
		expired:    expired,
	}
}

func (e *FuncEffect[T]) GetEffectKey() any {
	return e.key()
}

func (e *FuncEffect[T]) ActivateEffect(target T) {
	e.activate(target)
}

func (e *FuncEffect[T]) DeactivateEffect(target T) {
	e.deactivate(target)
}

func (e *FuncEffect[T]) IsExpired() bool {
	return e.expired()
}

type EffectMap[T any] struct {
	target  T
	effects map[any]Effect[T]
}

func (m *EffectMap[T]) AddEffect(effect Effect[T]) {
	key := effect.GetEffectKey()
	existingEffect := m.GetEffect(key)
	if (existingEffect != nil) {
		m.RemoveEffect(existingEffect)
	}
	effect.ActivateEffect(m.target)
	m.effects[key] = effect
}

func (m *EffectMap[T]) RemoveEffect(effect Effect[T]) {
	effect.DeactivateEffect(m.target)
	key := effect.GetEffectKey()
	delete(m.effects, key)
}

func (m *EffectMap[T]) GetEffect(key any) Effect[T] {
	return m.effects[key]
}
