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
	datetimeUnitEra             = 'G'
	datetimeUnitYear            = 'y'
	datetimeUnitMonth           = 'M'
	datetimeUnitDayOfWeek       = 'E'
	datetimeUnitDay             = 'd'
	datetimeUnitHour12          = 'h'
	datetimeUnitHour24          = 'H'
	datetimeUnitMinute          = 'm'
	datetimeUnitSecond          = 's'
	datetimeUnitPeriod          = 'a'
	datetimeUnitPeriodQuarter   = 'b'
	datetimeUnitPeriodRange     = 'B'
	datetimeUnitQuarter         = 'Q'
	datetimeUnitTimeZone        = 'z'
	datetimeUnitTimeZoneISO8601 = 'Z'
)

// The sequence length of datetime unit characters indicates how they should be rendered.
const (
	datetimeFormatLength1Plus       = 1
	datetimeFormatLength2Plus       = 2
	datetimeFormatLengthAbbreviated = 3
	datetimeFormatLengthWide        = 4
	datetimeFormatLengthNarrow      = 5
	datetimeFormatLengthShort       = 6
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
			result += mapTimePattern(pattern.Pattern)
		}
	}

	return result
}

func mapTimePattern(pattern string) string {
	switch pattern[0:1] {
	case string(datetimeUnitYear):
		return convertTimeYear(pattern)
	case string(datetimeUnitMonth):
		return convertTimeMonth(pattern)
	case string(datetimeUnitDayOfWeek):
		return convertTimeDayOfWeek(pattern)
	case string(datetimeUnitDay):
		return convertTimeDay(pattern)
	case string(datetimeUnitHour12):
		return convertTimeHour12(pattern)
	case string(datetimeUnitHour24):
		return "15"
	case string(datetimeUnitMinute):
		return convertTimeMinute(pattern)
	case string(datetimeUnitSecond):
		return convertTimeSecond(pattern)
	case string(datetimeUnitPeriod):
		return "PM"
	case string(datetimeUnitTimeZone):
		return "MST"
	case string(datetimeUnitTimeZoneISO8601):
		return convertTimeZoneISO8601(pattern)
	default:
		// ignore unmappable tokens
		return ""
	}
}

func convertTimeYear(pattern string) string {
	switch len(pattern) {
	case datetimeFormatLength2Plus:
		return "06"
	default:
		return "2006"
	}
}

func convertTimeMonth(pattern string) string {
	switch len(pattern) {
	case datetimeFormatLength1Plus:
		return "1"
	case datetimeFormatLength2Plus:
		return "01"
	case datetimeFormatLengthAbbreviated:
		return "Jan"
	case datetimeFormatLengthWide:
		return "January"
	default:
		return "January"
	}
}

func convertTimeDayOfWeek(pattern string) string {
	switch len(pattern) {
	case datetimeFormatLength1Plus:
		return "Mon"
	default:
		return "Monday"
	}
}

func convertTimeDay(pattern string) string {
	switch len(pattern) {
	case datetimeFormatLength1Plus:
		return "2"
	default:
		return "02"
	}
}

func convertTimeHour12(pattern string) string {
	switch len(pattern) {
	case datetimeFormatLength1Plus:
		return "3"
	default:
		return "03"
	}
}

func convertTimeMinute(pattern string) string {
	switch len(pattern) {
	case datetimeFormatLength1Plus:
		return "4"
	default:
		return "04"
	}
}

func convertTimeSecond(pattern string) string {
	switch len(pattern) {
	case datetimeFormatLength1Plus:
		return "5"
	default:
		return "05"
	}
}

func convertTimeZoneISO8601(pattern string) string {
	switch len(pattern) {
	case datetimeFormatLength1Plus, datetimeFormatLength2Plus, datetimeFormatLengthAbbreviated:
		return "-0700"
	case datetimeFormatLengthNarrow:
		return "-07:00:00"
	default:
		return "Z0700"
	}
}
