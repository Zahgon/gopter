package gen

import (
	"reflect"

	"github.com/leanovate/gopter"
)

type sliceShrinkOne struct {
	original      reflect.Value
	index         int
	elementShrink gopter.Shrink
}

func (s *sliceShrinkOne) Next() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

// SliceShrinkerOne creates a slice shrinker from a shrinker for the elements of the slice.
// The length of the slice will remains unchanged, instead each element is shrunk after the
// other.
func SliceShrinkerOne(elementShrinker gopter.Shrinker) gopter.Shrinker {
	_ = "STUB: not implemented"
	return *new(gopter.Shrinker)
}

type sliceShrink struct {
	original    reflect.Value
	length      int
	offset      int
	chunkLength int
}

func (s *sliceShrink) Next() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

// SliceShrinker creates a slice shrinker from a shrinker for the elements of the slice.
// The length of the slice will be shrunk as well
func SliceShrinker(elementShrinker gopter.Shrinker) gopter.Shrinker {
	_ = "STUB: not implemented"
	return *new(gopter.Shrinker)
}
