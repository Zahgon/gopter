package commands

import (
	"github.com/leanovate/gopter"
)

// Replay a sequence of commands on a system for regression testing
func Replay(systemUnderTest SystemUnderTest, initialState State, commands ...Command) *gopter.PropResult {
	_ = "STUB: not implemented"
	return nil
}
