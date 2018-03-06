package flaggy

import (
	"fmt"
	"os"
)

// SubCommand represents a subcommand which contains a set of child
// subcommands along with a set of flags relevant to it.  Parsing
// runs until a subcommand is detected by matching its name and
// position.  Once a matching subcommand is found, the next set
// of parsing occurs within that matched subcommand.
type SubCommand struct {
	LongName        string
	ShortName       string
	Description     string
	Position        int // the position of this subcommand, not including flags
	SubCommands     []*SubCommand
	StringFlags     []*StringFlag
	IntFlags        []*IntFlag
	BoolFlags       []*BoolFlag
	PositionalFlags []*PositionalValue // order matters here
}

// NewSubCommand creates a new subcommand that can have flags or PositionalFlags
// added to it.  The position starts with 1, not 0
func NewSubCommand(relativeDepth int) *SubCommand {
	if relativeDepth < 0 {
		fmt.Println("Flaggy: Position of flags and positional arguments must never be below 1")
		os.Exit(2)
	}
	return &SubCommand{}
}

// Parse causes the argument parser to parse based on the os.Args []string.
// depth specifies the non-flag subcommand positional depth
func (sc *SubCommand) parse(depth int) error {

	// TODO - determine string flags
	// TODO - determine int flags
	// TODO - determine bool flags
	// TODO - exclude normal flags in --key=value, -key=value, --key value,
	//        and -key value format before continuing
	// TODO - determine positional value flags by positional value
	// TODO - will parsing positionals before subcommands lead to positionals
	//        being parsed that shouldnt be?
	// TODO - determine subcommands and parse them by positional value ane name

	var err error

	// parse this subcommand's flags out of the command
	for i, v := range os.Args {
		argType := determineArgType(v)
		switch argType {
		case ArgIsPositional:
			debugPrint("Arg is positional:", v)
			// TODO - parse for positional arg or subcommand taking depth into account
		case ArgIsFlagWithSpace:
			debugPrint("Arg is flag with space:", v, i)
			// TODO - parse next arg as value to this flag
		case ArgIsFlagWithValue:
			debugPrint("Arg is flag with value:", v)
			// TODO - parse flag into key and value
			key, value := parseArgWithValue(v)
			debugPrint("Parsed key", key, "to value", value)
		}
	}

	// parse all child subcommand's flags
	for _, child := range sc.SubCommands {
		err = child.parse(depth + 1) // more depth for the next subcommand
		if err != nil {
			return err
		}
	}

	return nil
}

// AddSubcommand adds a possible subcommand to the ArgumentParser.
func (sc *SubCommand) AddSubcommand(newSC *SubCommand) {
	sc.SubCommands = append(sc.SubCommands, newSC)
}

// AddStringFlag adds a new string flag
func (sc *SubCommand) AddStringFlag(assignmentVar *string, shortName string, longName string, description string) {
	newStringFlag := StringFlag{}
	newStringFlag.AssignmentVar = assignmentVar
	newStringFlag.ShortName = shortName
	newStringFlag.LongName = longName
	newStringFlag.Description = description
	sc.StringFlags = append(sc.StringFlags, &newStringFlag)
}

// AddBoolFlag adds a new bool flag
func (sc *SubCommand) AddBoolFlag(assignmentVar *bool, shortName string, longName string, description string) {
	newBoolFlag := BoolFlag{}
	newBoolFlag.AssignmentVar = assignmentVar
	newBoolFlag.ShortName = shortName
	newBoolFlag.LongName = longName
	newBoolFlag.Description = description
	sc.BoolFlags = append(sc.BoolFlags, &newBoolFlag)
}

// AddIntFlag adds a new int flag
func (sc *SubCommand) AddIntFlag(assignmentVar *int, shortName string, longName string, description string) {
	newIntFlag := IntFlag{}
	newIntFlag.AssignmentVar = assignmentVar
	newIntFlag.ShortName = shortName
	newIntFlag.LongName = longName
	newIntFlag.Description = description
	sc.IntFlags = append(sc.IntFlags, &newIntFlag)
}
