package core

type DieFormula struct {
	n int
	d *Die
	c int
}

func (df *DieFormula) Roll() int {
	rolls := 0
	for range df.n {
		rolls += df.d.Roll()
	}
	return rolls + df.c
}
