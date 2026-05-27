package gen

import (
	"reflect"

	"github.com/leanovate/gopter"
)

// SliceOf generates an arbitrary slice of generated elements
// genParams.MaxSize sets an (exclusive) upper limit on the size of the slice
// genParams.MinSize sets an (inclusive) lower limit on the size of the slice
func SliceOf(elementGen gopter.Gen, typeOverrides ...reflect.Type) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}

// SliceOfN generates a slice of generated elements with definied length
func SliceOfN(desiredlen int, elementGen gopter.Gen, typeOverrides ...reflect.Type) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}

func genSlice(elementGen gopter.Gen, genParams *gopter.GenParameters, desiredlen int, typeOverride reflect.Type) (reflect.Value, func(interface{}) bool, gopter.Shrinker) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil, *new(gopter.Shrinker)
}

func forAllSieve(elementSieve func(interface{}) bool) func(interface{}) bool {
	_ = "STUB: not implemented"
	return nil
}
