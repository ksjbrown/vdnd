package core

type Effect[T any] interface {
	ApplyTo(target T)
}

