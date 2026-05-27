package gopter

// PropArg contains information about the specific values for a certain property check.
// This is mostly used for reporting when a property has falsified.
type PropArg struct {
	Arg              interface{}
	ArgFormatted     string
	OrigArg          interface{}
	OrigArgFormatted string
	Label            string
	Shrinks          int
}

func (p *PropArg) String() string { _ = "STUB: not implemented"; return "" }

// PropArgs is a list of PropArg.
type PropArgs []*PropArg

// NewPropArg creates a new PropArg.
func NewPropArg(genResult *GenResult, shrinks int, value interface{}, valueFormated string, origValue interface{}, origValueFormated string) *PropArg {
	_ = "STUB: not implemented"
	return nil
}
