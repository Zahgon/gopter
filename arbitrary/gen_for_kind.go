package arbitrary

import (
	"reflect"

	"github.com/leanovate/gopter"
)

func mapBoolish(to reflect.Type, v interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func mapIntish(to reflect.Type, v interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func mapUintish(to reflect.Type, v interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func mapFloatish(to reflect.Type, v interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func mapComplexish(to reflect.Type, v interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func mapStringish(to reflect.Type, v interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (a *Arbitraries) genForKind(rt reflect.Type) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}
