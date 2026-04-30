package core

// Action represents some action that a character can perform.
//
// The action can be applied to a certain encounter, in which the
// action is considered performed.
type Action struct {
	id   ActionID
	kind ActionKind
}

func (a *Action) ID() ActionID {
	return a.id
}

func (a *Action) Kind() ActionKind {
	return a.kind
}

type ActionID uint

type ActionKind uint
