package flaggy

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Subcommand represents a subcommand which contains a set of child
// subcommands along with a set of flags relevant to it.  Parsing
// runs until a subcommand is detected by matching its name and
// position.  Once a matching subcommand is found, the next set
// of parsing occurs within that matched subcommand.
type Subcommand struct {
	Name            string
	Description     string
	Position        int // the position of this subcommand, not including flags
	Subcommands     []*Subcommand
	StringFlags     []*StringFlag
	IntFlags        []*IntFlag
	BoolFlags       []*BoolFlag
	PositionalFlags []*PositionalValue
	SubcommandUsed  bool // indicates this subcommand was found and parsed
}

// NewSubcommand creates a new subcommand that can have flags or PositionalFlags
// added to it.  The position starts with 1, not 0
func NewSubcommand(name string, relativeDepth int) *Subcommand {
	if relativeDepth < 0 {
		fmt.Println("Flaggy: Position of flags and positional arguments must never be below 1")
		os.Exit(2)
	}
	return &Subcommand{
		Name:     name,
		Position: relativeDepth,
	}
}

// Parse causes the argument parser to parse based on the os.Args []string.
// depth specifies the non-flag subcommand positional depth
func (sc *Subcommand) parse(p *Parser, args []string, depth int) error {

	debugPrint("Parsing for depth of", depth)

	// if a command is parsed, its used
	sc.SubcommandUsed = true

	// holds positional arguments after flags are parsed out of the arguments
	positionalOnlyArguments := []string{}

	// indicates we should skip the next argument, like when parsing a flag
	// that seperates key and value by space
	var skipNext bool

	// endArgfound indicates that a -- was found and everything
	// remaining should be added to the trailing arguments slices
	var endArgFound bool

	// find all the normal flags (not positional) and parse them out
	for i, a := range args {

		if endArgFound {
			p.TrailingArguments = append(p.TrailingArguments, a)
			continue
		}

		// skip this run if specified
		if skipNext {
			skipNext = false
			continue
		}

		// determine what kind of flag this is
		argType := determineArgType(a)

		// depending on the flag type, parse the key and value out, then apply it
		switch argType {
		case ArgIsFinal:
			debugPrint("Arg is final:", a)
			endArgFound = true
		case ArgIsPositional:
			// debugPrint("Arg is positional:", a)
			// Add this positional argument into a slice of their own, so that
			// we can determine if its a subcommand or positional value later
			positionalOnlyArguments = append(positionalOnlyArguments, a)
		case ArgIsFlagWithSpace:
			debugPrint("Arg is flag with space:", a, i)
			// parse next arg as value to this flag and apply to subcommand flags
			skipNext = true
			err := sc.setValueForKey(a, os.Args[i+1])
			if err != nil {
				return err
			}
			continue
		case ArgIsFlagWithValue:
			debugPrint("Arg is flag with value:", a)
			// parse flag into key and value and apply to subcommand flags
			key, val := parseArgWithValue(a)
			err := sc.setValueForKey(key, val)
			if err != nil {
				return err
			}
			debugPrint("Parsed key", key, "to value", val)
		}
	}

	// loop over positional values and look for their matching positional
	// parameter, or their positional command.  If neither are found, then
	// we throw an error
	for pos, v := range positionalOnlyArguments {

		// the first relative positional argument will be human natural at position 1
		// but offset for the depth of relative commands being parsed for currently.
		relativeDepth := (1 + pos) - depth
		if relativeDepth < 1 {
			// debugPrint("relativeDepth", relativeDepth, "<", 1, "skipping")
			continue
		}

		debugPrint("Parsing positional only position", pos)
		var foundPositionalMatch bool
		// determine positional value flags by positional value and depth of parser

		// determine subcommands and parse them by positional value and name
		// TODO - will parsing positionals before subcommands lead to positionals
		//        being parsed that shouldnt be?
		for _, cmd := range sc.Subcommands {
			// debugPrint(relativeDepth, "==", cmd.Position, "and", v, "==", cmd.Name)
			if relativeDepth == cmd.Position && (v == cmd.Name) {
				debugPrint("Parsing positional subcommand", cmd.Name, "at relativeDepth", relativeDepth)
				return cmd.parse(p, args, depth+1) // continue recursive positional parsing
			}
		}

		// determine positional args  and parse them by positional value and name
		for _, val := range sc.PositionalFlags {
			if relativeDepth == val.Position {
				debugPrint("Found a positional value at relativePos:", relativeDepth, "value:", v)
				// defrerence the struct pointer, then set the pointer property within it
				*val.AssignmentVar = v
				debugPrint("set positional to value", *val.AssignmentVar)
				foundPositionalMatch = true
			}
		}

		// if no match found for this arg, throw an error
		if !foundPositionalMatch {
			// if no positional match was detected, then we throw an error because this
			// argument is unexpected.
			return errors.New("Unable to find a positional subcommand or argument after " + sc.Name + " at relative pos: " + strconv.Itoa(relativeDepth) + " value was: " + v)
		}

	}

	return nil
}

// AddSubcommand adds a possible subcommand to the Parser.
func (sc *Subcommand) AddSubcommand(newSC *Subcommand) error {

	// ensure no subcommands at this depth with this name
	for _, other := range sc.Subcommands {
		if newSC.Position == other.Position {
			return errors.New("Unable to add subcommand because one already exists at position: " + strconv.Itoa(newSC.Position))
		}
	}

	// ensure no positionals at this depth
	for _, other := range sc.PositionalFlags {
		if newSC.Position == other.Position {
			return errors.New("Unable to add subcommand because a positional value already exists at position: " + strconv.Itoa(newSC.Position))
		}
	}

	sc.Subcommands = append(sc.Subcommands, newSC)

	return nil
}

// AddStringFlag adds a new string flag
func (sc *Subcommand) AddStringFlag(assignmentVar *string, shortName string, longName string, description string) {
	newStringFlag := StringFlag{}
	newStringFlag.AssignmentVar = assignmentVar
	newStringFlag.ShortName = shortName
	newStringFlag.LongName = longName
	newStringFlag.Description = description
	sc.StringFlags = append(sc.StringFlags, &newStringFlag)
}

// AddBoolFlag adds a new bool flag
func (sc *Subcommand) AddBoolFlag(assignmentVar *bool, shortName string, longName string, description string) {
	newBoolFlag := BoolFlag{}
	newBoolFlag.AssignmentVar = assignmentVar
	newBoolFlag.ShortName = shortName
	newBoolFlag.LongName = longName
	newBoolFlag.Description = description
	sc.BoolFlags = append(sc.BoolFlags, &newBoolFlag)
}

// AddIntFlag adds a new int flag
func (sc *Subcommand) AddIntFlag(assignmentVar *int, shortName string, longName string, description string) {
	newIntFlag := IntFlag{}
	newIntFlag.AssignmentVar = assignmentVar
	newIntFlag.ShortName = shortName
	newIntFlag.LongName = longName
	newIntFlag.Description = description
	sc.IntFlags = append(sc.IntFlags, &newIntFlag)
}

// AddPositionalValue adds a positional value to the subcommand.  the
// relativePosition starts at 1 and is relative to the subcommand it belongs to
func (sc *Subcommand) AddPositionalValue(relativePosition int, assignmentVar *string, name string, description string) error {

	// ensure no other positionals are at this depth
	for _, other := range sc.PositionalFlags {
		if relativePosition == other.Position {
			return errors.New("Unable to add positional value because one already exists at position: " + strconv.Itoa(relativePosition))
		}
	}

	// ensure no subcommands at this depth
	for _, other := range sc.Subcommands {
		if relativePosition == other.Position {
			return errors.New("Unable to add positional value a subcommand already exists at position: " + strconv.Itoa(relativePosition))
		}
	}

	newPositionalValue := PositionalValue{
		Name:          name,
		Position:      relativePosition,
		AssignmentVar: assignmentVar,
		Description:   description,
	}
	sc.PositionalFlags = append(sc.PositionalFlags, &newPositionalValue)

	return nil
}

// SetValueForKey sets the value for the specified key. If setting a bool
// value, then leave the value field empty (``)
func (sc *Subcommand) setValueForKey(key string, value string) error {

	debugPrint("setting value for", key, "to", value)

	// check for and assign string flags
	for _, f := range sc.StringFlags {
		if f.ShortName == key || f.LongName == key {
			newValue := value
			f.AssignmentVar = &newValue
			debugPrint("Set string flag with key ", key, "to value", value)
		}
	}

	// check for and assign int flags
	for _, f := range sc.IntFlags {
		if f.ShortName == key || f.LongName == key {
			newValue, err := strconv.Atoi(value)
			if err != nil {
				return errors.New("Unable to convert flag to int: " + err.Error())
			}
			f.AssignmentVar = &newValue
			debugPrint("Set int flag with key ", key, "to value", value)
		}
	}

	// check for and assign bool flags
	for _, f := range sc.BoolFlags {
		if f.ShortName == key || f.LongName == key {

			// if there is no value specified, we assume the bool flag to be toggled on
			if value == `` {
				newValue := true
				f.AssignmentVar = &newValue
				return nil
			}
			newValue, err := strconv.ParseBool(value)
			if err != nil {
				return errors.New("Unable to convert flag to bool: " + err.Error())
			}
			f.AssignmentVar = &newValue
			debugPrint("Set bool flag with key ", key, "to value", value)
		}
	}

	return errors.New("An unexpected flag was specified: " + key)
}
