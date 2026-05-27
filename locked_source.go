// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Taken from golang lockedSource implementation https://github.com/golang/go/blob/master/src/math/rand/rand.go#L371-L410

package gopter

import (
	"math/rand"
	"sync"
)

type lockedSource struct {
	lk  sync.Mutex
	src rand.Source64
}

// NewLockedSource takes a seed and returns a new
// lockedSource for use with rand.New
func NewLockedSource(seed int64) *lockedSource { _ = "STUB: not implemented"; return nil }

func (r *lockedSource) Int63() (n int64) { _ = "STUB: not implemented"; return 0 }

func (r *lockedSource) Uint64() (n uint64) { _ = "STUB: not implemented"; return 0 }

func (r *lockedSource) Seed(seed int64) { _ = "STUB: not implemented"; return }

// seedPos implements Seed for a lockedSource without a race condition.
func (r *lockedSource) seedPos(seed int64, readPos *int8) { _ = "STUB: not implemented"; return }

// read implements Read for a lockedSource without a race condition.
func (r *lockedSource) read(p []byte, readVal *int64, readPos *int8) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func read(p []byte, int63 func() int64, readVal *int64, readPos *int8) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
