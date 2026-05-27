package gen

import (
	"reflect"

	"github.com/leanovate/gopter"
)

// MapOf generates an arbitrary map of generated kay values.
// genParams.MaxSize sets an (exclusive) upper limit on the size of the map
// genParams.MinSize sets an (inclusive) lower limit on the size of the map
func MapOf(keyGen, elementGen gopter.Gen) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}

func genMap(keyGen, elementGen gopter.Gen, genParams *gopter.GenParameters, len int) (reflect.Value, func(interface{}) bool, gopter.Shrinker, func(interface{}) bool, gopter.Shrinker) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil, *new(gopter.Shrinker), nil, *new(gopter.Shrinker)
}

func forAllKeyValueSieve(keySieve, elementSieve func(interface{}) bool) func(interface{}) bool {
	_ = "STUB: not implemented"
	return nil
}
