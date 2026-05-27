package gopter

// Prop represent some kind of property that (drums please) can and should be checked
type Prop func(*GenParameters) *PropResult

// SaveProp creates s save property by handling all panics from an inner property
func SaveProp(prop Prop) Prop { _ = "STUB: not implemented"; return *new(Prop) }

// Check the property using specific parameters
func (prop Prop) Check(parameters *TestParameters) *TestResult {
	_ = "STUB: not implemented"
	return nil
}
