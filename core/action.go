package core

import "encoding/json"

type ActionData struct {
	Kind ActionKind
	Data json.RawMessage
}

type Action interface {
	Kind() ActionKind
}

type ActionID uint

type ActionKind uint
