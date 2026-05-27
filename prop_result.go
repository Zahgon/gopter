package gopter

type propStatus int

const (
	// PropProof THe property was proved (i.e. it is known to be correct and will be always true)
	PropProof propStatus = iota
	// PropTrue The property was true this time
	PropTrue
	// PropFalse The property was false this time
	PropFalse
	// PropUndecided The property has no clear outcome this time
	PropUndecided
	// PropError The property has generated an error
	PropError
)

func (s propStatus) String() string { _ = "STUB: not implemented"; return "" }

// PropResult contains the result of a property
type PropResult struct {
	Status     propStatus
	Error      error
	ErrorStack []byte
	Args       []*PropArg
	Labels     []string
}

// NewPropResult create a PropResult with label
func NewPropResult(success bool, label string) *PropResult { _ = "STUB: not implemented"; return nil }

// Success checks if the result was successful
func (r *PropResult) Success() bool { _ = "STUB: not implemented"; return false }

// WithArgs sets argument descriptors to the PropResult for reporting
func (r *PropResult) WithArgs(args []*PropArg) *PropResult { _ = "STUB: not implemented"; return nil }

// AddArgs add argument descriptors to the PropResult for reporting
func (r *PropResult) AddArgs(args ...*PropArg) *PropResult { _ = "STUB: not implemented"; return nil }

// And combines two PropResult by an and operation.
// The resulting PropResult will be only true if both PropResults are true.
func (r *PropResult) And(other *PropResult) *PropResult { _ = "STUB: not implemented"; return nil }

func (r *PropResult) mergeWith(other *PropResult, status propStatus) *PropResult {
	_ = "STUB: not implemented"
	return nil
}
