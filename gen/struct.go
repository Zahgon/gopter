package gen

import (
	"reflect"

	"github.com/leanovate/gopter"
)

// Struct generates a given struct type.
// rt has to be the reflect type of the struct, gens contains a map of field generators.
// Note that the result types of the generators in gen have to match the type of the corresponding
// field in the struct. Also note that only public fields of a struct can be generated
func Struct(rt reflect.Type, gens map[string]gopter.Gen) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}

// StructPtr generates pointers to a given struct type.
// Note that StructPtr does not generate nil, if you want to include nil in your
// testing you should combine gen.PtrOf with gen.Struct.
// rt has to be the reflect type of the struct, gens contains a map of field generators.
// Note that the result types of the generators in gen have to match the type of the corresponding
// field in the struct. Also note that only public fields of a struct can be generated
func StructPtr(rt reflect.Type, gens map[string]gopter.Gen) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}

// checkFieldsMatch panics unless the keys in gens exactly match the public
// fields on rt. With an extra bool argument of value "true", it only panics if
// there's a key in gens which is not a field on rt.
func checkFieldsMatch(
	rt reflect.Type,
	gens map[string]gopter.Gen,
	allowFieldsWithNoGenerator ...bool,
) {
	_ = "STUB: not implemented"
	return
}

// Don't check that every field is present in gens

// Check that every field is present in gens

// StrictStruct behaves the same as Struct, except it requires the keys in gens
// to exactly match the public fields of rt. It panics if gens contains extra
// keys, or has missing keys.
//
// If given a third true argument, it only requires the keys of gens to be
// fields of rt. In that case, unspecified fields will remain unset.
func StrictStruct(
	rt reflect.Type,
	gens map[string]gopter.Gen,
	allowFieldsWithNoGenerator ...bool,
) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}

// StrictStructPtr behaves the same as StructPtr, except it requires the keys in
// gens to exactly match the public fields of rt. It panics if gens contains
// extra keys, or has missing keys.
//
// If given a third true argument, it only requires the keys of gens to be
// fields of rt. In that case, unspecified fields will remain unset.
func StrictStructPtr(
	rt reflect.Type,
	gens map[string]gopter.Gen,
	allowFieldsWithNoGenerator ...bool,
) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}
