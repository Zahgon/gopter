package gen

import (
	"time"

	"github.com/leanovate/gopter"
)

// Time generates an arbitrary time.Time within year [0, 9999]
func Time() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Ensure year in [0, 9999]

// AnyTime generates an arbitrary time.Time struct (might be way out of bounds of any reason)
func AnyTime() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// TimeRange generates an arbitrary time.Time with a range
// from defines the start of the time range
// duration defines the overall duration of the time range
func TimeRange(from time.Time, duration time.Duration) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}
