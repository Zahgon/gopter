package gen

import (
	"github.com/leanovate/gopter"
)

type nilShrink struct {
	done bool
}

func (s *nilShrink) Next() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

// PtrShrinker convert a value shrinker to a pointer to value shrinker
func PtrShrinker(elementShrinker gopter.Shrinker) gopter.Shrinker {
	_ = "STUB: not implemented"
	return *new(gopter.Shrinker)
}
