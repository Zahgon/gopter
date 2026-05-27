package gopter

import "testing"

// Properties is a collection of properties that should be checked in a test
type Properties struct {
	parameters *TestParameters
	props      map[string]Prop
	propNames  []string
}

// NewProperties create new Properties with given test parameters.
// If parameters is nil default test parameters will be used
func NewProperties(parameters *TestParameters) *Properties { _ = "STUB: not implemented"; return nil }

// Property add/defines a property in a test.
func (p *Properties) Property(name string, prop Prop) { _ = "STUB: not implemented"; return }

// Run checks all definied propertiesand reports the result
func (p *Properties) Run(reporter Reporter) bool { _ = "STUB: not implemented"; return false }

// TestingRun checks all definied properties with a testing.T context.
// This the preferred wait to run property tests as part of a go unit test.
func (p *Properties) TestingRun(t *testing.T, opts ...interface{}) {
	_ = "STUB: not implemented"
	return
}
