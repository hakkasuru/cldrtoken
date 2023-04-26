package parser

import (
	"errors"
	"strings"
)

// literals
const (
	datetimeTimeSeparator = ':'
	datetimeLiteral       = '\''
)

// CLDRToken holder for each pattern and their types
type CLDRToken struct {
	Pattern   string
	IsLiteral bool
}

// CLDRParser for cldr tokens
type CLDRParser struct{}

// Parse takes in cldr format string tokens and returns a sequence of cldr tokens
func (p CLDRParser) Parse(tokens string) ([]CLDRToken, error) {
	patterns := []CLDRToken{}
	tokenLen := len(tokens)

	for i := 0; i < tokenLen; {
		char := tokens[i : i+1]

		if char == string(datetimeLiteral) {
			// find the next single quote
			// create a literal out of everything between the quotes
			// and set i to the position after the second quote

			if i == tokenLen-1 {
				return nil, errors.New("malformed datetime format")
			}

			nextQuote := strings.Index(tokens[i+1:], string(datetimeLiteral))
			if nextQuote == -1 {
				return nil, errors.New("malformed datetime format")
			}

			cldrToken := CLDRToken{
				Pattern:   tokens[i+1 : nextQuote+i+1],
				IsLiteral: true,
			}

			patterns = append(patterns, cldrToken)
			i = nextQuote + i + 2

			continue
		}

		if (char >= "a" && char <= "z") || (char >= "A" && char <= "Z") {
			// this represents a format unit
			// find the entire sequence of the same character
			endChar := lastSequenceIndex(tokens[i:]) + i

			cldrToken := CLDRToken{
				Pattern:   tokens[i : endChar+1],
				IsLiteral: false,
			}

			patterns = append(patterns, cldrToken)
			i = endChar + 1

			continue
		}

		if char == string(datetimeTimeSeparator) {
			cldrToken := CLDRToken{
				Pattern:   ":",
				IsLiteral: true,
			}

			patterns = append(patterns, cldrToken)
			i++

			continue
		}

		cldrToken := CLDRToken{
			Pattern:   char,
			IsLiteral: true,
		}

		patterns = append(patterns, cldrToken)
		i++

		continue
	}

	return patterns, nil
}

/*
lastSequenceIndex looks at the first character in a string and returns the
last digits of the first sequence of that character. For example:
- ABC: 0
- AAB: 1
- ABA: 0
- AAA: 2
*/
func lastSequenceIndex(str string) int {
	if len(str) == 0 {
		return -1
	}

	if len(str) == 1 {
		return 0
	}

	sequenceChar := str[0:1]
	lastPos := 0
	for i := 1; i < len(str); i++ {
		if str[i:i+1] != sequenceChar {
			break
		}

		lastPos = i
	}

	return lastPos
}
