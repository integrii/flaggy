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
	Used            bool // indicates this subcommand was found and parsed
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

// parseAllFlagsFromArgs parses the non-positional flags such as -f or -v=value
// out of the supplied args and returns the positional items in order
func (sc *Subcommand) parseAllFlagsFromArgs(p *Parser, args []string) ([]string, error) {

	var err error
	var positionalOnlyArguments []string

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
			debugPrint("Arg", i, "is final:", a)
			endArgFound = true
		case ArgIsPositional:
			// debugPrint("Arg is positional:", a)
			// Add this positional argument into a slice of their own, so that
			// we can determine if its a subcommand or positional value later
			positionalOnlyArguments = append(positionalOnlyArguments, a)
		case ArgIsFlagWithSpace:
			debugPrint("Arg", i, "is flag with space:", a, i)
			// parse next arg as value to this flag and apply to subcommand flags
			// if the flag is a bool flag, then we check for a following positional
			if flagIsBool(sc, p, a) {
				switch {
				case args[i+1] == "true":
					err = setValueForParsers(a, "true", p, sc)
					if err != nil {
						return []string{}, err
					}
					skipNext = true
				case args[i+1] == "false":
					err = setValueForParsers(a, "false", p, sc)
					if err != nil {
						return []string{}, err
					}
					skipNext = true
				default:
					// if the next value was not true or false, we assume this bool
					// flag stands alone and should be assumed to mean true.  In this
					// case, we do not skip the next flag in the argument list.
					err = setValueForParsers(a, "true", p, sc)
					if err != nil {
						return []string{}, err
					}
				}
				// by default, we just assign the next argument to the value and continue
				if err != nil {
					return []string{}, err
				}
				continue
			}
			err = setValueForParsers(a, args[i+1], p, sc)
			if err != nil {
				return []string{}, err
			}
		case ArgIsFlagWithValue:
			debugPrint("Arg", i, "is flag with value:", a)
			// parse flag into key and value and apply to subcommand flags
			key, val := parseArgWithValue(a)
			err = setValueForParsers(key, val, p, sc)
			if err != nil {
				return []string{}, err
			}
			debugPrint("Parsed key", key, "to value", val)
		}
	}

	return positionalOnlyArguments, nil
}

// Parse causes the argument parser to parse based on the supplied []string.
// depth specifies the non-flag subcommand positional depth
func (sc *Subcommand) parse(p *Parser, args []string, depth int) error {

	debugPrint("Parsing for depth of", depth)

	// if a command is parsed, its used
	sc.Used = true

	// Parse the normal flags out of the argument list and retain the positionals.
	// Apply the flags to the parent parser and the current subcommand context.
	positionalOnlyArguments, err := sc.parseAllFlagsFromArgs(p, args)
	if err != nil {
		return err
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

		debugPrint("Parsing positional only position", pos, "with value", v)
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
			}
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
// value, then send "true" or "false" as strings.  The returned bool indicates
// that a value was set.
func (sc *Subcommand) SetValueForKey(key string, value string) (bool, error) {

	// check for and assign string flags
	for _, f := range sc.StringFlags {
		if f.ShortName == key || f.LongName == key {
			debugPrint("Setting string value for", key, "to", value)
			newValue := value
			f.AssignmentVar = &newValue
			debugPrint("Set string flag with key ", key, "to value", value)
			return true, nil
		}
	}

	// check for and assign int flags
	for _, f := range sc.IntFlags {
		if f.ShortName == key || f.LongName == key {
			debugPrint("Setting int value for", key, "to", value)
			newValue, err := strconv.Atoi(value)
			if err != nil {
				return false, errors.New("Unable to convert flag to int: " + err.Error())
			}
			f.AssignmentVar = &newValue
			debugPrint("Set int flag with key ", key, "to value", value)
			return true, nil
		}
	}

	// check for and assign bool flags
	for _, f := range sc.BoolFlags {
		if f.ShortName == key || f.LongName == key {
			debugPrint("Setting bool value for", key, "to", value)
			newValue, err := strconv.ParseBool(value)
			if err != nil {
				return false, errors.New("Unable to convert flag to bool: " + err.Error())
			}
			f.AssignmentVar = &newValue
			debugPrint("Set bool flag with key ", key, "to value", value)
			return true, nil
		}
	}

	return false, nil
}
