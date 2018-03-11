package flaggy

import (
	"fmt"
	"strings"
)

// Flag holds the base methods for all flag types
type Flag struct {
	ShortName   string
	LongName    string
	Description string
}

// StringFlag represents a flag that is converted into a string value.
type StringFlag struct {
	Flag
	AssignmentVar *string
}

// IntFlag represents a flag that is converted into an int value.
type IntFlag struct {
	Flag
	AssignmentVar *int
}

// BoolFlag represents a flag that is converted into a bool value.
type BoolFlag struct {
	Flag
	AssignmentVar *bool
}

const ArgIsPositional = "positional"       // subcommand or positional value
const ArgIsFlagWithSpace = "flagWithSpace" // -f path or --file path
const ArgIsFlagWithValue = "flagWithValue" // -f=path or --file=path
const ArgIsFinal = "final"                 // the final argument only '--'

// determineArgType determines if the specified arg is a flag with space
// separated value, a flag with a connected value, or neither (positional)
func determineArgType(arg string) string {

	// if the arg is --, then its the final arg
	if arg == "--" {
		return ArgIsFinal
	}

	// if it has the prefix --, then its a long flag
	if strings.HasPrefix(arg, "--") {
		// if it contains an equals, it is a joined value
		if strings.Contains(arg, "=") {
			return ArgIsFlagWithValue
		}
		return ArgIsFlagWithSpace
	}

	// if it has the prefix -, then its a short flag
	if strings.HasPrefix(arg, "-") {
		// if it contains an equals, it is a joined value
		if strings.Contains(arg, "=") {
			return ArgIsFlagWithValue
		}
		return ArgIsFlagWithSpace
	}

	return ArgIsPositional
}

// parseArgWithValue parses a key=value concatentated argument
func parseArgWithValue(arg string) (key string, value string) {

	// remove up to two minuses from start of flag
	arg = strings.TrimPrefix(arg, "-")
	arg = strings.TrimPrefix(arg, "-")

	debugPrint("parseArgWithValue", arg)

	// break at the equals
	args := strings.SplitN(arg, "=", 2)

	// if its a bool arg, with no explicit value, we return a blank
	if len(args) == 1 {
		return args[0], ""
	}

	// if its a key and value pair, we return those
	if len(args) == 2 {
		return args[0], args[1]
	}

	fmt.Println("Warning: attempted to parseArgWithValue but did not have correct parameter count.", arg, "->", args)
	return "", ""
}

// parseFlagToName parses a flag with space value down to a key name:
//     --path -> path
//     -p -> p
func parseFlagToName(arg string) string {
	// remove minus from start
	arg = strings.TrimLeft(arg, "-")
	arg = strings.TrimLeft(arg, "-")
	return arg
}

// flagIsBool determines if the flag is a bool within the specified parser
// and subcommand's context
func flagIsBool(sc *Subcommand, p *Parser, key string) bool {
	for _, f := range sc.BoolFlags {
		if f.ShortName == key || f.LongName == key {
			return true
		}
	}
	for _, f := range p.BoolFlags {
		if f.ShortName == key || f.LongName == key {
			return true
		}
	}

	// by default, the answer is false
	return false
}
