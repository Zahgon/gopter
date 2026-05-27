package gen

import (
	"github.com/leanovate/gopter"
)

// OneConstOf generate one of a list of constant values
func OneConstOf(consts ...interface{}) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}

// OneGenOf generate one value from a a list of generators
func OneGenOf(gens ...gopter.Gen) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }
