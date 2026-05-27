package gen

import (
	"regexp/syntax"

	"github.com/leanovate/gopter"
)

// RegexMatch generates matches for a given regular expression
// regexStr is supposed to conform to the perl regular expression syntax
func RegexMatch(regexStr string) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

func regexMatchGen(regex *syntax.Regexp) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}

func runeToString(v rune) string { _ = "STUB: not implemented"; return "" }
