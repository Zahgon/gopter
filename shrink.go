package gopter

// Shrink is a stream of shrunk down values.
// Once the result of a shrink is false, it is considered to be exhausted.
// Important notes for implementors:
//   - Ensure that the returned stream is finite, even though shrinking will
//     eventually be aborted, infinite streams may result in very slow running
//     test.
//   - Ensure that modifications to the returned value will not affect the
//     internal state of your Shrink. If in doubt return by value not by reference
type Shrink func() (interface{}, bool)

// Filter creates a shrink filtered by a condition
func (s Shrink) Filter(condition func(interface{}) bool) Shrink {
	_ = "STUB: not implemented"
	return *new(Shrink)
}

// Map creates a shrink by applying a converter to each element of a shrink.
// f: has to be a function with one parameter (matching the generated value) and a single return.
func (s Shrink) Map(f interface{}) Shrink { _ = "STUB: not implemented"; return *new(Shrink) }

// All collects all shrinks as a slice. Use with care as this might create
// large results depending on the complexity of the shrink
func (s Shrink) All() []interface{} { _ = "STUB: not implemented"; return nil }

type concatedShrink struct {
	index   int
	shrinks []Shrink
}

func (c *concatedShrink) Next() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

// ConcatShrinks concats an array of shrinks to a single shrinks
func ConcatShrinks(shrinks ...Shrink) Shrink { _ = "STUB: not implemented"; return *new(Shrink) }

type interleaved struct {
	first          Shrink
	second         Shrink
	firstExhausted bool
	secondExhaused bool
	state          bool
}

func (i *interleaved) Next() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

// Interleave this shrink with another
// Both shrinks are expected to produce the same result
func (s Shrink) Interleave(other Shrink) Shrink { _ = "STUB: not implemented"; return *new(Shrink) }

// Shrinker creates a shrink for a given value
type Shrinker func(value interface{}) Shrink

type elementShrink struct {
	original      []interface{}
	index         int
	elementShrink Shrink
}

func (e *elementShrink) Next() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

// CombineShrinker create a shrinker by combining a list of shrinkers.
// The resulting shrinker will shrink an []interface{} where each element will be shrunk by
// the corresonding shrinker in 'shrinkers'.
// This method is implicitly used by CombineGens.
func CombineShrinker(shrinkers ...Shrinker) Shrinker {
	_ = "STUB: not implemented"
	return *new(Shrinker)
}

// NoShrink is an empty shrink.
var NoShrink = Shrink(func() (interface{}, bool) {
	return nil, false
})

// NoShrinker is a shrinker for NoShrink, i.e. a Shrinker that will not shrink any values.
// This is the default Shrinker if none is provided.
var NoShrinker = Shrinker(func(value interface{}) Shrink {
	return NoShrink
})
