package gen

import (
	"reflect"

	"github.com/leanovate/gopter"
)

type mapShrinkOne struct {
	original         reflect.Value
	key              reflect.Value
	keyShrink        gopter.Shrink
	elementShrink    gopter.Shrink
	state            bool
	keyExhausted     bool
	lastKey          interface{}
	elementExhausted bool
	lastElement      interface{}
}

func (s *mapShrinkOne) nextKeyValue() (interface{}, interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

func (s *mapShrinkOne) Next() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

// MapShrinkerOne creates a map shrinker from a shrinker for the key values of a map.
// The length of the map will remain (mostly) unchanged, instead each key value pair is
// shrunk after the other.
func MapShrinkerOne(keyShrinker, elementShrinker gopter.Shrinker) gopter.Shrinker {
	_ = "STUB: not implemented"
	return *new(gopter.Shrinker)
}

type mapShrink struct {
	original     reflect.Value
	originalKeys []reflect.Value
	length       int
	offset       int
	chunkLength  int
}

func (s *mapShrink) Next() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

// MapShrinker creates a map shrinker from shrinker for the key values.
// The length of the map will be shrunk as well
func MapShrinker(keyShrinker, elementShrinker gopter.Shrinker) gopter.Shrinker {
	_ = "STUB: not implemented"
	return *new(gopter.Shrinker)
}
