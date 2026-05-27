package gopter

import (
	"math/rand"
)

// TestParameters to run property tests
type TestParameters struct {
	MinSuccessfulTests int
	// MinSize is an (inclusive) lower limit on the size of the parameters
	MinSize int
	// MaxSize is an (exclusive) upper limit on the size of the parameters
	MaxSize         int
	MaxShrinkCount  int
	seed            int64
	Rng             *rand.Rand
	Workers         int
	MaxDiscardRatio float64
}

func (t *TestParameters) Seed() int64 { _ = "STUB: not implemented"; return 0 }

func (t *TestParameters) SetSeed(seed int64) { _ = "STUB: not implemented"; return }

// DefaultTestParameterWithSeeds creates reasonable default Parameters for most cases based on a fixed RNG-seed
func DefaultTestParametersWithSeed(seed int64) *TestParameters {
	_ = "STUB: not implemented"
	return nil
}

// DefaultTestParameterWithSeeds creates reasonable default Parameters for most cases with an undefined RNG-seed
func DefaultTestParameters() *TestParameters { _ = "STUB: not implemented"; return nil }
