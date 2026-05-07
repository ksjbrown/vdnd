package core

type DeltaMutator interface {
	GetDelta() int
}

// -- Implementations --

type DeltaMutatorMap[T comparable] struct {
	mutators map[T]DeltaMutator
}

func (m *DeltaMutatorMap[T]) SetDeltaMutator(key T, mutator DeltaMutator) {
	if m.mutators == nil {
		m.mutators = make(map[T]DeltaMutator)
	}
	m.mutators[key] = mutator
}

func (m *DeltaMutatorMap[T]) AppendDeltaMutator(key T, mutator DeltaMutator) {
	mutator := m.GetDeltaMutator(key)
	compositeMutator, ok := mutator.(*DeltaMutatorComposite) {

	}
}

func (m *DeltaMutatorMap[T]) GetDeltaMutator(key T) DeltaMutator {
	return m.mutators[key]
}

func (m *DeltaMutatorMap[T]) GetDelta(key T) int {
	mutator := m.GetDeltaMutator(key)
	if (mutator == nil) {
		return 0
	}
	return mutator.GetDelta()
}

type DeltaMutatorComposite struct {
	mutators []DeltaMutator
}

func (m *DeltaMutatorComposite) GetDelta() int {
	delta := 0
	for _, mutator := range m.mutators {
		delta += mutator.GetDelta()
	}
	return delta
}

func (m *DeltaMutatorComposite) AddDeltaMutator(mutator DeltaMutator) {
	m.mutators = append(m.mutators, mutator)
}

type DeltaMutatorFunc struct {
	fn func() int
}

func (m *DeltaMutatorFunc) GetDelta() int {
	return m.fn()
}
