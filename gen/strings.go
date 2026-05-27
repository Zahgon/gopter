package gen

import (
	"unicode"

	"github.com/leanovate/gopter"
)

// RuneRange generates runes within a given range
func RuneRange(min, max rune) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Rune generates an arbitrary character rune
func Rune() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// RuneNoControl generates an arbitrary character rune that is not a control character
func RuneNoControl() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

func genRune(int64Gen gopter.Gen) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// NumChar generates arbitrary numberic character runes
func NumChar() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// AlphaUpperChar generates arbitrary uppercase alpha character runes
func AlphaUpperChar() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// AlphaLowerChar generates arbitrary lowercase alpha character runes
func AlphaLowerChar() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// AlphaChar generates arbitrary character runes (upper- and lowercase)
func AlphaChar() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// AlphaNumChar generates arbitrary alpha-numeric character runes
func AlphaNumChar() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// UnicodeChar generates arbitrary character runes with a given unicode table
func UnicodeChar(table *unicode.RangeTable) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}

// AnyString generates an arbitrary string
func AnyString() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// AlphaString generates an arbitrary string with letters
func AlphaString() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// NumString generates an arbitrary string with digits
func NumString() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Identifier generates an arbitrary identifier string
// Identitiers are supporsed to start with a lowercase letter and contain only
// letters and digits
func Identifier() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// UnicodeString generates an arbitrary string from a given
// unicode table.
func UnicodeString(table *unicode.RangeTable) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}

func genString(runeGen gopter.Gen, runeSieve func(ch rune) bool) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}

func runesToString(v []rune) string { _ = "STUB: not implemented"; return "" }
