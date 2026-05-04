package core

type Health struct {
	value int
	max   int
	temp  int
}

func (h *Health) Value() int {
	return h.value
}
func (h *Health) Max() int {
	return h.max
}
func (h *Health) Temp() int {
	return h.temp
}
