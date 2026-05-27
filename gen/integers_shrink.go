package gen

import (
	"github.com/leanovate/gopter"
)

type int64Shrink struct {
	original int64
	half     int64
}

func (s *int64Shrink) Next() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

type uint64Shrink struct {
	original uint64
	half     uint64
}

func (s *uint64Shrink) Next() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

// Int64Shrinker is a shrinker for int64 numbers
func Int64Shrinker(v interface{}) gopter.Shrink {
	_ = "STUB: not implemented"
	return *new(gopter.Shrink)
}

// UInt64Shrinker is a shrinker for uint64 numbers
func UInt64Shrinker(v interface{}) gopter.Shrink {
	_ = "STUB: not implemented"
	return *new(gopter.Shrink)
}

// Int32Shrinker is a shrinker for int32 numbers
func Int32Shrinker(v interface{}) gopter.Shrink {
	_ = "STUB: not implemented"
	return *new(gopter.Shrink)
}

// UInt32Shrinker is a shrinker for uint32 numbers
func UInt32Shrinker(v interface{}) gopter.Shrink {
	_ = "STUB: not implemented"
	return *new(gopter.Shrink)
}

// Int16Shrinker is a shrinker for int16 numbers
func Int16Shrinker(v interface{}) gopter.Shrink {
	_ = "STUB: not implemented"
	return *new(gopter.Shrink)
}

// UInt16Shrinker is a shrinker for uint16 numbers
func UInt16Shrinker(v interface{}) gopter.Shrink {
	_ = "STUB: not implemented"
	return *new(gopter.Shrink)
}

// Int8Shrinker is a shrinker for int8 numbers
func Int8Shrinker(v interface{}) gopter.Shrink {
	_ = "STUB: not implemented"
	return *new(gopter.Shrink)
}

// UInt8Shrinker is a shrinker for uint8 numbers
func UInt8Shrinker(v interface{}) gopter.Shrink {
	_ = "STUB: not implemented"
	return *new(gopter.Shrink)
}

// IntShrinker is a shrinker for int numbers
func IntShrinker(v interface{}) gopter.Shrink {
	_ = "STUB: not implemented"
	return *new(gopter.Shrink)
}

// UIntShrinker is a shrinker for uint numbers
func UIntShrinker(v interface{}) gopter.Shrink {
	_ = "STUB: not implemented"
	return *new(gopter.Shrink)
}

// int64RangeShrinker returns a shrinker that shrinks toward min, staying within [min, max].
func int64RangeShrinker(min, max int64) gopter.Shrinker {
	_ = "STUB: not implemented"
	return *new(gopter.Shrinker)
}

// Shrink toward min by halving the distance

// Convert from distance-from-min back to actual value
