package gen

import (
	"reflect"

	"github.com/leanovate/gopter"
)

type arrayShrinkOne struct {
	original      reflect.Value
	index         int
	elementShrink gopter.Shrink
}

func (s *arrayShrinkOne) Next() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

// ArrayShrinkerOne creates an array shrinker from a shrinker for the elements of the slice.
// The length of the array will remains unchanged, instead each element is shrunk after the
// other.
func ArrayShrinkerOne(elementShrinker gopter.Shrinker) gopter.Shrinker {
	_ = "STUB: not implemented"
	return *new(gopter.Shrinker)
}
