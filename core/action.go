package core

// Action represents some action that a character can perform.
//
// The action can be applied to a certain encounter, in which the
// action is considered performed.
//
// Still need to plan, but will probably 
type Action struct {
	ID   ActionID
	kind ActionKind
}

func (a *Action) Kind() ActionKind {
	return a.kind
}

type ActionID uint

type ActionKind uint
