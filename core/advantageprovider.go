package core

type AdvantageProvider interface {
	GetAdvantage() AdvantageKind
}
