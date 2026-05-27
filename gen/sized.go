package gen

import (
	"github.com/leanovate/gopter"
)

// Sized derives a generator from based on size
// This honors the `MinSize` and `MaxSize` of the `GenParameters` of the test suite.
// Keep an eye on memory consumption, by default MaxSize is 100.
func Sized(f func(int) gopter.Gen) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }
