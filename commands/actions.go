package commands

import (
	"github.com/leanovate/gopter"
)

type shrinkableCommand struct {
	command      Command
	commandSieve func(v interface{}) bool
	shrinker     gopter.Shrinker
}

func (s shrinkableCommand) shrink() gopter.Shrink {
	_ = "STUB: not implemented"
	return *new(gopter.Shrink)
}

func (s shrinkableCommand) String() string { _ = "STUB: not implemented"; return "" }

type actions struct {
	// initialStateProvider has to reset/recreate the initial state exactly the
	// same every time.
	initialStateProvider func() State
	sequentialCommands   []shrinkableCommand
	// parallel commands will come later
}

func (a *actions) String() string { _ = "STUB: not implemented"; return "" }

func (a *actions) run(systemUnderTest SystemUnderTest) *gopter.PropResult {
	_ = "STUB: not implemented"
	return nil
}

type sizedCommands struct {
	state    State
	commands []shrinkableCommand
}

func actionsShrinker(v interface{}) gopter.Shrink {
	_ = "STUB: not implemented"
	return *new(gopter.Shrink)
}

func genActions(commands Commands) gopter.Gen { _ = "STUB: not implemented"; return *new(gopter.Gen) }

func genSizedCommands(commands Commands, initialStateProvider func() State) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}
