package flaggy

import "fmt"

// Help represents the values needed to render a Help page
type Help struct {
	Subcommands    []HelpSubcommand
	Positionals    []HelpPositional
	Flags          []HelpFlag
	UsageString    string
	CommandName    string
	PrependMessage string
	AppendMessage  string
	Message        string
	Description    string
}

// HelpSubcommand is used to template subcommand Help output
type HelpSubcommand struct {
	ShortName   string
	LongName    string
	Description string
	Position    int
}

// HelpPositional is used to template positional Help output
type HelpPositional struct {
	Name         string
	Description  string
	Required     bool
	Position     int
	DefaultValue string
}

// HelpFlag is used to template string flag Help output
type HelpFlag struct {
	ShortName    string
	LongName     string
	Description  string
	DefaultValue string
}

// ExtractValues extracts Help template values from a subcommand
func (h *Help) ExtractValues(sc *Subcommand, message string) {
	// extract Help values
	// prependMessage string
	h.PrependMessage = sc.AdditionalHelpPrepend
	// appendMessage  string
	h.AppendMessage = sc.AdditionalHelpAppend
	// message string
	h.Message = message
	// command name
	h.CommandName = sc.Name
	// description
	h.Description = sc.Description
	// subcommands    []HelpSubcommand
	for _, cmd := range sc.Subcommands {
		if cmd.Hidden {
			continue
		}
		newHelpSubcommand := HelpSubcommand{
			ShortName:   cmd.ShortName,
			LongName:    cmd.Name,
			Description: cmd.Description,
			Position:    cmd.Position,
		}
		h.Subcommands = append(h.Subcommands, newHelpSubcommand)
	}
	// positionals    []HelpPositional
	for _, pos := range sc.PositionalFlags {
		if pos.Hidden {
			continue
		}
		newHelpPositional := HelpPositional{
			Name:         pos.Name,
			Position:     pos.Position,
			Description:  pos.Description,
			Required:     pos.Required,
			DefaultValue: *pos.AssignmentVar,
		}
		h.Positionals = append(h.Positionals, newHelpPositional)
	}

	for _, f := range sc.Flags {
		if f.Hidden {
			continue
		}
		defaultValue, err := f.returnAssignmentVarValueAsString()
		if err != nil {
			fmt.Println("Error when generating help template values:", err)
		}
		if defaultValue == "<nil>" {
			defaultValue = ""
		}

		newHelpFlag := HelpFlag{
			ShortName:    f.ShortName,
			LongName:     f.LongName,
			Description:  f.Description,
			DefaultValue: defaultValue,
		}
		h.Flags = append(h.Flags, newHelpFlag)
	}
	// formulate the usage string
	// first, we capture all the command and positional names by position
	commandsByPosition := make(map[int]string)
	for _, pos := range sc.PositionalFlags {
		if pos.Hidden {
			continue
		}
		if len(commandsByPosition[pos.Position]) > 0 {
			commandsByPosition[pos.Position] = commandsByPosition[pos.Position] + "|" + pos.Name
		} else {
			commandsByPosition[pos.Position] = pos.Name
		}
	}
	for _, cmd := range sc.Subcommands {
		if cmd.Hidden {
			continue
		}
		if len(commandsByPosition[cmd.Position]) > 0 {
			commandsByPosition[cmd.Position] = commandsByPosition[cmd.Position] + "|" + cmd.Name
		} else {
			commandsByPosition[cmd.Position] = cmd.Name
		}
	}

	// find the highest position count in the map
	var highestPosition int
	for i := range commandsByPosition {
		if i > highestPosition {
			highestPosition = i
		}
	}

	// find each positional value and make our final string
	var usageString = sc.Name
	for i := 1; i <= highestPosition; i++ {
		if len(commandsByPosition[i]) > 0 {
			usageString = usageString + " [" + commandsByPosition[i] + "]"
		} else {
			// dont keep listing after the first position without any properties
			// it will be impossible to reach anything beyond here anyway
			break
		}
	}

	h.UsageString = usageString

}
