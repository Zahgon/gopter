package gen

import (
	"github.com/leanovate/gopter"
)

// Float64Range generates float64 numbers within a given range
func Float64Range(min, max float64) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Float64 generates arbitrary float64 numbers that do not contain NaN or Inf
func Float64() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Float32Range generates float32 numbers within a given range
func Float32Range(min, max float32) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Float32 generates arbitrary float32 numbers that do not contain NaN or Inf
func Float32() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }
