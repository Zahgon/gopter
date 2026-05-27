package gopter

import (
	"reflect"
)

// BiMapper is a bi-directional (or bijective) mapper of a tuple of values (up)
// to another tuple of values (down).
type BiMapper struct {
	UpTypes    []reflect.Type
	DownTypes  []reflect.Type
	Downstream reflect.Value
	Upstream   reflect.Value
}

// NewBiMapper creates a BiMapper of two functions `downstream` and its
// inverse `upstream`.
// That is: The return values of `downstream` must match the parameters of
// `upstream` and vice versa.
func NewBiMapper(downstream interface{}, upstream interface{}) *BiMapper {
	_ = "STUB: not implemented"
	return nil
}

// ConvertUp calls the Upstream function on the arguments in the down array
// and returns the results.
func (b *BiMapper) ConvertUp(down []interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ConvertDown calls the Downstream function on the elements of the up array
// and returns the results.
func (b *BiMapper) ConvertDown(up []interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}
