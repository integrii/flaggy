package flaggy

// Help represents the values needed to render a Help page
type Help struct {
	Subcommands    []HelpSubcommand
	Positionals    []HelpPositional
	StringFlags    []HelpFlag
	IntFlags       []HelpFlag
	BoolFlags      []HelpFlag
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
	Name        string
	Description string
	Required    bool
	Position    int
}

// HelpFlag is used to template string flag Help output
type HelpFlag struct {
	ShortName   string
	LongName    string
	Description string
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
	for _, sc := range sc.Subcommands {
		newHelpSubcommand := HelpSubcommand{
			ShortName:   sc.ShortName,
			LongName:    sc.Name,
			Description: sc.Description,
			Position:    sc.Position,
		}
		h.Subcommands = append(h.Subcommands, newHelpSubcommand)
	}
	// positionals    []HelpPositional
	for _, pos := range sc.PositionalFlags {
		newHelpPositional := HelpPositional{
			Name:        pos.Name,
			Position:    pos.Position,
			Description: pos.Description,
			Required:    pos.Required,
		}
		h.Positionals = append(h.Positionals, newHelpPositional)
	}

	// flags          []HelpFlag
	for _, f := range sc.StringFlags {
		newHelpFlag := HelpFlag{
			ShortName:   f.ShortName,
			LongName:    f.LongName,
			Description: f.Description,
		}
		h.StringFlags = append(h.StringFlags, newHelpFlag)
	}
	for _, f := range sc.IntFlags {
		newHelpFlag := HelpFlag{
			ShortName:   f.ShortName,
			LongName:    f.LongName,
			Description: f.Description,
		}
		h.IntFlags = append(h.IntFlags, newHelpFlag)
	}
	for _, f := range sc.BoolFlags {
		newHelpFlag := HelpFlag{
			ShortName:   f.ShortName,
			LongName:    f.LongName,
			Description: f.Description,
		}
		h.BoolFlags = append(h.BoolFlags, newHelpFlag)
	}
}
