package gen

import "github.com/leanovate/gopter"

// Complex128Box generate complex128 numbers within a rectangle/box in the complex plane
func Complex128Box(min, max complex128) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}

// Complex128 generate arbitrary complex128 numbers
func Complex128() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Complex64Box generate complex64 numbers within a rectangle/box in the complex plane
func Complex64Box(min, max complex64) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}

// Complex64 generate arbitrary complex64 numbers
func Complex64() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }
