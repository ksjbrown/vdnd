package core

type D20Test struct {
	Die           *Die
	Bonus int
	AdvantageKind AdvantageKind
	Target 		  int
}

func NewD20Test(die *Die, bonus int, advantage AdvantageKind, target int) *D20Test {
	return &D20Test{
		Die: die,
		Bonus: bonus,
		AdvantageKind: advantage,
		Target: target,
	}
}

func (t *D20Test) DoTest() *D20TestResult {
	const Advantage = D20TestAdvantageKindAdvantage
	const Disadvantage = D20TestAdvantageKindDisadvantage

	result := &D20TestResult{}
	roll := t.Die.Roll()

	switch t.AdvantageKind {
	case Advantage:
		result.Roll = max(roll, t.Die.Roll())
	case Disadvantage:
		result.Roll = min(roll, t.Die.Roll())
	default:
		result.Roll = roll
	}

	result.IsSuccess = result.Roll >= t.Target
	return result
}

type AdvantageKind int

const (
	AdvantageKindNone         AdvantageKind = 0x01
	D20TestAdvantageKindAdvantage    AdvantageKind = 0x02
	D20TestAdvantageKindDisadvantage AdvantageKind = 0x03
)

type D20TestResult struct {
	Roll      int
	IsSuccess bool
}
