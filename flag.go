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

const ArgIsPositional = 1    // subcommand or positional value
const ArgIsFlagWithSpace = 2 // -f path or --file path
const ArgIsFlagWithValue = 3 // -f=path or --file=path

// determineArgType determines if the specified arg is a flag with space
// separated value, a flag with a connected value, or neither (positional)
func determineArgType(arg string) int {
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
	// remove minus from start
	arg = strings.TrimLeft(arg, "-")
	arg = strings.TrimLeft(arg, "-")

	// break at the equals
	args := strings.SplitN(arg, "=", 1)
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
