package gopter

import "reflect"

// GenResult contains the result of a generator.
type GenResult struct {
	Labels     []string
	Shrinker   Shrinker
	ResultType reflect.Type
	Result     interface{}
	Sieve      func(interface{}) bool
}

// NewGenResult creates a new generator result from for a concrete value and
// shrinker.
// Note: The concrete value "result" not be nil
func NewGenResult(result interface{}, shrinker Shrinker) *GenResult {
	_ = "STUB: not implemented"
	return nil
}

// NewEmptyResult creates an empty generator result.
// Unless the sieve does not explicitly allow it, empty (i.e. nil-valued)
// results are considered invalid.
func NewEmptyResult(resultType reflect.Type) *GenResult { _ = "STUB: not implemented"; return nil }

// Retrieve gets the concrete generator result.
// If the result is invalid or does not pass the sieve there is no concrete
// value and the property using the generator should be undecided.
func (r *GenResult) Retrieve() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

// RetrieveAsValue get the concrete generator result as reflect value.
// If the result is invalid or does not pass the sieve there is no concrete
// value and the property using the generator should be undecided.
func (r *GenResult) RetrieveAsValue() (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}
