package gopter

import (
	"math/rand"
)

// GenParameters encapsulates the parameters for all generators.
type GenParameters struct {
	MinSize        int
	MaxSize        int
	MaxShrinkCount int
	Rng            *rand.Rand
}

// WithSize modifies the size parameter. The size parameter defines an upper bound for the size of
// generated slices or strings.
func (p *GenParameters) WithSize(size int) *GenParameters { _ = "STUB: not implemented"; return nil }

// NextBool create a random boolean using the underlying Rng.
func (p *GenParameters) NextBool() bool { _ = "STUB: not implemented"; return false }

// NextInt64 create a random int64 using the underlying Rng.
func (p *GenParameters) NextInt64() int64 { _ = "STUB: not implemented"; return 0 }

// NextUint64 create a random uint64 using the underlying Rng.
func (p *GenParameters) NextUint64() uint64 { _ = "STUB: not implemented"; return 0 }

// CloneWithSeed clone the current parameters with a new seed.
// This is useful to create subsections that can rerun (provided you keep the
// seed)
func (p *GenParameters) CloneWithSeed(seed int64) *GenParameters {
	_ = "STUB: not implemented"
	return nil
}

// DefaultGenParameters creates default GenParameters.
func DefaultGenParameters() *GenParameters { _ = "STUB: not implemented"; return nil }

// MinGenParameters creates minimal GenParameters.
// Note: Most likely you do not want to use these for actual testing
func MinGenParameters() *GenParameters { _ = "STUB: not implemented"; return nil }
