package core

type D20Test int

func NewD20Test(advantage AdvantageKind, bonus int) D20Test {
	die := NewDieD20()
	roll := die.Roll()

	switch advantage {
	case AdvantageKindAdvantage:
		roll = max(roll, die.Roll())
	case AdvantageKindDisadvantage:
		roll = min(roll, die.Roll())
	}

	value := roll + bonus
	return D20Test(value)
}

type AdvantageKind int

const (
	AdvantageKindNone         AdvantageKind = 0x01
	AdvantageKindAdvantage    AdvantageKind = 0x02
	AdvantageKindDisadvantage AdvantageKind = 0x03
)
