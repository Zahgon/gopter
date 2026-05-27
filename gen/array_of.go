package gen

import (
	"reflect"

	"github.com/leanovate/gopter"
)

// ArrayOfN generates an array of generated elements with definied length
func ArrayOfN(desiredlen int, elementGen gopter.Gen, typeOverrides ...reflect.Type) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}

func genArray(elementGen gopter.Gen, genParams *gopter.GenParameters, desiredlen int, typeOverride reflect.Type) (reflect.Value, func(interface{}) bool, gopter.Shrinker) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil, *new(gopter.Shrinker)
}
