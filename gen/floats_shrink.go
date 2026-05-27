package gen

import (
	"github.com/leanovate/gopter"
)

type float64Shrink struct {
	original float64
	half     float64
}

func (s *float64Shrink) isZeroOrVeryClose() bool { _ = "STUB: not implemented"; return false }

func (s *float64Shrink) Next() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

// Float64Shrinker is a shrinker for float64 numbers
func Float64Shrinker(v interface{}) gopter.Shrink {
	_ = "STUB: not implemented"
	return *new(gopter.Shrink)
}

// Float32Shrinker is a shrinker for float32 numbers
func Float32Shrinker(v interface{}) gopter.Shrink {
	_ = "STUB: not implemented"
	return *new(gopter.Shrink)
}
