package gen

import (
	"github.com/leanovate/gopter"
)

// PickN takes a slice and a number, and returns a generator that produces
// subsets of exactly n items in their original order, randomly selected without duplicates
// Example: PickN([]int{1,2,3,4,5}, 3) might produce []int{1,3,5} or []int{2,4,5}
// Note: This generator has no shrinker. Use small slices to keep failing test feedback readable.
func PickN(items interface{}, number int) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}
