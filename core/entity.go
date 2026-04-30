package core

// What can be done to an Entity?
// It can be:
// - inspected
// - interacted
// - attacked
// These may all be able to be handled under a single ApplyAction
// Or maybe we keep it simple and just allow modifying fields directly.

// Entity represents any inanimate or animate "thing" in the universe.
type Entity struct {
	ID EntityID
	Name string
	HP Health
	Effects []Effect
}

type EntityID uint

