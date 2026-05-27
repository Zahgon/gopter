package prop

import (
	"github.com/leanovate/gopter"
)

/*
ForAllNoShrink creates a property that requires the check condition to be true for all values.
As the name suggests the generated values will not be shrunk if the condition falsiies.

"condition" has to be a function with the same number of parameters as the provided
generators "gens". The function may return a simple bool (true means that the
condition has passed), a string (empty string means that condition has passed),
a *PropResult, or one of former combined with an error.
*/
func ForAllNoShrink(condition interface{}, gens ...gopter.Gen) gopter.Prop {
	_ = "STUB: not implemented"
	return *new(gopter.Prop)
}

// ForAllNoShrink1 creates a property that requires the check condition to be true for all values
// As the name suggests the generated values will not be shrunk if the condition falsiies
func ForAllNoShrink1(gen gopter.Gen, check func(interface{}) (interface{}, error)) gopter.Prop {
	_ = "STUB: not implemented"
	return *new(gopter.Prop)
}
