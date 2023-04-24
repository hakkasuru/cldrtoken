package core

// Core converter interface
type Core interface {
	// Convert cldr tokens to golang tokens, returns empty string something goes wrong
	Convert(token string) string
}

type noopCore struct{}

// NewNoopCore no operation core
func NewNoopCore() Core                        { return noopCore{} }
func (c noopCore) Convert(token string) string { return "" }
