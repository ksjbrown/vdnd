package core

import "math/rand"

type IntProvider interface {
	GetInt(min, max int) int
}

type RandomIntProvider struct {}

func (r *RandomIntProvider) GetInt(min, max int) int {
	return min + rand.Intn(max)
}

