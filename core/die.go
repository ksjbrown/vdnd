package core

import "math/rand"

// Die represents an n-sided die
type Die int

func (d Die) Roll() int {
	return rand.Intn(int(d + 1))
}

type DieFormula struct {
	n int
	d Die
	c int
}

func (df *DieFormula) Roll() int {
	rolls := 0
	for range df.n {
		rolls += df.d.Roll()
	}
	return rolls + df.c
}
