package core

var globalDieIntProvider IntProvider

func init() {
	SetGlobalDieIntProvider(&RandomIntProvider{})
}

func SetGlobalDieIntProvider(ip IntProvider) {
	globalDieIntProvider = ip
}

// Die represents an n-sided die
type Die struct {
	sides int
	ip    IntProvider
}

func NewDie(sides int, ip IntProvider) *Die {
	return &Die{
		sides: sides,
		ip:    ip,
	}
}

func (d *Die) Roll() int {
	return d.ip.GetInt(1, d.sides)
}

func NewDieD20() *Die {
	dip := globalDieIntProvider
	return NewDie(20, dip)
}
