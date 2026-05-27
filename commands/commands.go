package commands

import (
	"github.com/leanovate/gopter"
)

// Commands provide an entry point for testing a stateful system
type Commands interface {
	// NewSystemUnderTest should create a new/isolated system under test
	NewSystemUnderTest(initialState State) SystemUnderTest
	// DestroySystemUnderTest may perform any cleanup tasks to destroy a system
	DestroySystemUnderTest(SystemUnderTest)
	// GenInitialState provides a generator for the initial State.
	// IMPORTANT: The generated state itself may be mutable, but this generator
	// is supposed to generate a clean and reproductable state every time.
	// Do not use an external random generator and be especially vary about
	// `gen.Const(<pointer to some mutable struct>)`.
	GenInitialState() gopter.Gen
	// GenCommand provides a generator for applicable commands to for a state
	GenCommand(state State) gopter.Gen
	// InitialPreCondition checks if the initial state is valid
	InitialPreCondition(state State) bool
}

// ProtoCommands is a prototype implementation of the Commands interface
type ProtoCommands struct {
	NewSystemUnderTestFunc     func(initialState State) SystemUnderTest
	DestroySystemUnderTestFunc func(SystemUnderTest)
	InitialStateGen            gopter.Gen
	GenCommandFunc             func(State) gopter.Gen
	InitialPreConditionFunc    func(State) bool
}

// NewSystemUnderTest should create a new/isolated system under test
func (p *ProtoCommands) NewSystemUnderTest(initialState State) SystemUnderTest {
	_ = "STUB: not implemented"
	return *new(SystemUnderTest)
}

// DestroySystemUnderTest may perform any cleanup tasks to destroy a system
func (p *ProtoCommands) DestroySystemUnderTest(systemUnderTest SystemUnderTest) {
	_ = "STUB: not implemented"
	return
}

// GenCommand provides a generator for applicable commands to for a state
func (p *ProtoCommands) GenCommand(state State) gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}

// GenInitialState provides a generator for the initial State
func (p *ProtoCommands) GenInitialState() gopter.Gen {
	_ = "STUB: not implemented"
	return *new(gopter.Gen)
}

// InitialPreCondition checks if the initial state is valid
func (p *ProtoCommands) InitialPreCondition(state State) bool {
	_ = "STUB: not implemented"
	return false
}

// Prop creates a gopter.Prop from Commands
func Prop(commands Commands) gopter.Prop { _ = "STUB: not implemented"; return *new(gopter.Prop) }
