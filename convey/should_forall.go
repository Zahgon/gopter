package convey

// ShouldSucceedForAll checks that a check condition is be true for all values, if the
// condition falsiies the generated values will be shrunk.
//
// "condition" has to be a function with the same number of parameters as the provided
// generators "gens". The function may return a simple bool (true means that the
// condition has passed), a string (empty string means that condition has passed),
// a *PropResult, or one of former combined with an error.
func ShouldSucceedForAll(condition interface{}, params ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}
