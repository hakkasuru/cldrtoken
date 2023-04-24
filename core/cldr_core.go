package core

type CLDRCore struct{}

func NewCLDRCore() Core { return &CLDRCore{} }

func (c *CLDRCore) Convert(token string) string { return "" }
