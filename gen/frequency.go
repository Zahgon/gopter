package gen

import (
	"github.com/leanovate/gopter"
)

// Frequency combines multiple weighted generators of the the same result type
// The generators from weightedGens will be used accrding to the weight, i.e. generators
// with a hight weight will be used more often than generators with a low weight.
func Frequency(weightedGens map[int]gopter.Gen) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}
