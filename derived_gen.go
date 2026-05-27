package gopter

import (
	"reflect"
)

type derivedGen struct {
	biMapper   *BiMapper
	upGens     []Gen
	resultType reflect.Type
}

func (d *derivedGen) Generate(genParams *GenParameters) *GenResult {
	_ = "STUB: not implemented"
	return nil
}

func (d *derivedGen) Sieve(baseSieve ...func(interface{}) bool) func(interface{}) bool {
	_ = "STUB: not implemented"
	return nil
}

func (d *derivedGen) Shrinker(baseShrinker Shrinker) func(down interface{}) Shrink {
	_ = "STUB: not implemented"
	return nil
}

// DeriveGen derives a generator with shrinkers from a sequence of other
// generators mapped by a bijective function (BiMapper)
func DeriveGen(downstream interface{}, upstream interface{}, gens ...Gen) Gen {
	_ = "STUB: not implemented"
	return *new(Gen)
}
