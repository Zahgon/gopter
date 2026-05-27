package prop

import (
	"reflect"

	"github.com/leanovate/gopter"
)

func checkConditionFunc(check interface{}, numArgs int) (func([]reflect.Value) *gopter.PropResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
