package gen

import (
	"github.com/leanovate/gopter"
)

// PtrOf generates either a pointer to a generated element or a nil pointer
func PtrOf(elementGen gopter.Gen) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// To get the right pointer type we have to create a slice with one element
