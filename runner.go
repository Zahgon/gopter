package gopter

import (
	"sync"
)

type shouldStop func() bool

type worker func(int, shouldStop) *TestResult

type runner struct {
	sync.RWMutex
	parameters *TestParameters
	worker     worker
}

func (r *runner) mergeCheckResults(r1, r2 *TestResult) *TestResult {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) runWorkers() *TestResult { _ = "STUB: not implemented"; return nil }
