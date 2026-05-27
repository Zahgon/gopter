package gen

import (
	"github.com/leanovate/gopter"
)

// Int64Range generates int64 numbers within a given range
func Int64Range(min, max int64) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Check for range overflow

// UInt64Range generates uint64 numbers within a given range
func UInt64Range(min, max uint64) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Check overflow (i.e. max = MaxInt64, min = MinInt64)

// Int64 generates an arbitrary int64 number
func Int64() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// UInt64 generates an arbitrary Uint64 number
func UInt64() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Int32Range generates int32 numbers within a given range
func Int32Range(min, max int32) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// UInt32Range generates uint32 numbers within a given range
func UInt32Range(min, max uint32) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Int32 generate arbitrary int32 numbers
func Int32() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// UInt32 generate arbitrary int32 numbers
func UInt32() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Int16Range generates int16 numbers within a given range
func Int16Range(min, max int16) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// UInt16Range generates uint16 numbers within a given range
func UInt16Range(min, max uint16) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Int16 generate arbitrary int16 numbers
func Int16() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// UInt16 generate arbitrary uint16 numbers
func UInt16() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Int8Range generates int8 numbers within a given range
func Int8Range(min, max int8) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// UInt8Range generates uint8 numbers within a given range
func UInt8Range(min, max uint8) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Int8 generate arbitrary int8 numbers
func Int8() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// UInt8 generate arbitrary uint8 numbers
func UInt8() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// IntRange generates int numbers within a given range
func IntRange(min, max int) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Int generate arbitrary int numbers
func Int() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// UIntRange generates uint numbers within a given range
func UIntRange(min, max uint) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// UInt generate arbitrary uint numbers
func UInt() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

// Size just extracts the MaxSize field of the GenParameters.
// This can be helpful to generate limited integer value in a more structued
// manner.
func Size() gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

func int64To32(value int64) int32 { _ = "STUB: not implemented"; return 0 }

func uint64To32(value uint64) uint32 { _ = "STUB: not implemented"; return 0 }

func int64To16(value int64) int16 { _ = "STUB: not implemented"; return 0 }

func uint64To16(value uint64) uint16 { _ = "STUB: not implemented"; return 0 }

func int64To8(value int64) int8 { _ = "STUB: not implemented"; return 0 }

func uint64To8(value uint64) uint8 { _ = "STUB: not implemented"; return 0 }

func int64ToInt(value int64) int { _ = "STUB: not implemented"; return 0 }

func uint64ToUint(value uint64) uint { _ = "STUB: not implemented"; return 0 }
