package gen

import (
	"github.com/leanovate/gopter"
)

// WeightedGen adds a weight number to a generator.
// To be used as parameter to gen.Weighted
type WeightedGen struct {
	Weight int
	Gen    gopter.Gen
}

// Weighted combines multiple generators, where each generator has a weight.
// The weight of a generator is proportional to the probability that the
// generator gets selected.
func Weighted(weightedGens []WeightedGen) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}
