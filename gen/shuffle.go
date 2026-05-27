package gen

import (
	"github.com/leanovate/gopter"
)

// Shuffle takes a slice and returns a generator that produces shuffled copies
// Example: Shuffle([]int{1,2,3}) might produce []int{3,1,2} or []int{2,3,1}
// Note: This generator has no shrinker. Use small slices to keep failing test feedback readable.
func Shuffle(items interface{}) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }
