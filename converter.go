package cldrtoken

import "github.com/hakkasuru/cldrtoken/core"

type Converter struct {
	core core.Core
}

func NewNoopConverter() *Converter {
	return &Converter{}
}

func NewCLDRConverter() *Converter {
	return &Converter{}
}
