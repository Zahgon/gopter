package arbitrary

import (
	"github.com/leanovate/gopter"
)

/*
ForAll creates a property that requires the check condition to be true for all
values, if the condition falsiies the generated values will be shrunk.

"condition" has to be a function with the any number of parameters that can
generated in context of the Arbitraries. The function may return a simple bool,
a *PropResult, a boolean with error or a *PropResult with error.
*/
func (a *Arbitraries) ForAll(condition interface{}) gopter.Prop {
	_ = "STUB: not implemented"
	return *new(gopter.Prop)
}
