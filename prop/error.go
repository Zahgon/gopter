package prop

import "github.com/leanovate/gopter"

// ErrorProp creates a property that will always fail with an error.
// Mostly used as a fallback when setup/initialization fails
func ErrorProp(err error) gopter.Prop { _ = "STUB: not implemented"; return *new(gopter.Prop) }
