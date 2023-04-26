package core

import (
	"github.com/hakkasuru/cldrtoken/parser"
)

/*
Characters with special meaning in a datetime string:
Technically, all a-z,A-Z characters should be treated as if they represent a
datetime unit - but not all actually do. Any a-z,A-Z character that is
intended to be rendered as a literal a-z,A-Z character should be surrounded
by single quotes. There is currently no support for rendering a single quote
literal.
*/
const (
	datetimeUnitEra           = 'G'
	datetimeUnitYear          = 'y'
	datetimeUnitMonth         = 'M'
	datetimeUnitDayOfWeek     = 'E'
	datetimeUnitDay           = 'd'
	datetimeUnitHour12        = 'h'
	datetimeUnitHour24        = 'H'
	datetimeUnitMinute        = 'm'
	datetimeUnitSecond        = 's'
	datetimeUnitPeriod        = 'a'
	datetimeUnitPeriodQuarter = 'b'
	datetimeUnitPeriodRange   = 'B'
	datetimeUnitQuarter       = 'Q'
	datetimeUnitTimeZone1     = 'z'
	datetimeUnitTimeZone2     = 'v'
)

// The sequence length of datetime unit characters indicates how they should be rendered.
const (
	datetimeFormatLength1Plus       = 1
	datetimeFormatLength2Plus       = 2
	datetimeFormatLengthAbbreviated = 3
	datetimeFormatLengthWide        = 4
	datetimeFormatLengthNarrow      = 5
)

type CLDRCore struct {
	parser parser.CLDRParser
}

func NewCLDRCore() Core {
	return &CLDRCore{
		parser: parser.CLDRParser{},
	}
}

func (c *CLDRCore) Convert(token string) string {
	result := ""

	patterns, _ := c.parser.Parse(token)
	for _, pattern := range patterns {
		if pattern.IsLiteral {
			result += pattern.Pattern
		} else {

		}
	}

	return result
}

func mapDateTimePattern(pattern string) string {
	switch pattern[0:1] {
	case string(datetimeUnitEra):
		return ""
	case string(datetimeUnitYear):
		return ""
	case string(datetimeUnitMonth):
		return ""
	case string(datetimeUnitDayOfWeek):
		return ""
	case string(datetimeUnitDay):
		return ""
	case string(datetimeUnitHour12):
		return ""
	case string(datetimeUnitHour24):
		return ""
	case string(datetimeUnitMinute):
		return ""
	case string(datetimeUnitSecond):
		return ""
	case string(datetimeUnitPeriod):
		return ""
	case string(datetimeUnitPeriodQuarter):
		return ""
	case string(datetimeUnitPeriodRange):
		return ""
	case string(datetimeUnitQuarter):
		return ""
	case string(datetimeUnitTimeZone1):
		return ""
	case string(datetimeUnitTimeZone2):
		return ""
	default:
		return ""
	}
}
