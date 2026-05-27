package prop

import (
	"reflect"

	"github.com/leanovate/gopter"
)

var typeOfError = reflect.TypeOf((*error)(nil)).Elem()

/*
ForAll creates a property that requires the check condition to be true for all values, if the
condition falsiies the generated values will be shrunk.

"condition" has to be a function with the same number of parameters as the provided
generators "gens". The function may return a simple bool (true means that the
condition has passed), a string (empty string means that condition has passed),
a *PropResult, or one of former combined with an error.
*/
func ForAll(condition interface{}, gens ...gopter.Gen) gopter.Prop {
	_ = "STUB: not implemented"
	return *new(gopter.Prop)
}

// ForAll1 legacy interface to be removed in the future
func ForAll1(gen gopter.Gen, check func(v interface{}) (interface{}, error)) gopter.Prop {
	_ = "STUB: not implemented"
	return *new(gopter.Prop)
}

func shrinkValue(maxShrinkCount int, genResult *gopter.GenResult, origValue interface{}, orgiValueFormated string,
	firstFail *gopter.PropResult, check func(interface{}) *gopter.PropResult) (*gopter.PropResult, interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}

func firstFailure(shrink gopter.Shrink, check func(interface{}) *gopter.PropResult) (*gopter.PropResult, interface{}, string) {
	_ = "STUB: not implemented"
	return nil, nil, ""
}
