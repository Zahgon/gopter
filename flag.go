package gopter

// Flag is a convenient helper for an atomic boolean
type Flag struct {
	flag int32
}

// Get the value of the flag
func (f *Flag) Get() bool { _ = "STUB: not implemented"; return false }

// Set the the flag
func (f *Flag) Set() { _ = "STUB: not implemented"; return }

// Unset the flag
func (f *Flag) Unset() { _ = "STUB: not implemented"; return }
