package flaggy

// Help represents the values needed to render a Help page
type Help struct {
	subcommands    []HelpSubcommand
	positionals    []HelpPositional
	stringFlags    []HelpFlag
	intFlags       []HelpFlag
	boolFlags      []HelpFlag
	prependMessage string
	appendMessage  string
	message        string
}

// HelpSubcommand is used to template subcommand Help output
type HelpSubcommand struct {
	ShortName   string
	LongName    string
	Description string
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
	FlagType    string
}

// ExtractValues extracts Help template values from a subcommand
func (h *Help) ExtractValues(sc *Subcommand, message string) {
	// extract Help values
	// prependMessage string
	h.prependMessage = sc.AdditionalHelpPrepend
	// appendMessage  string
	h.appendMessage = sc.AdditionalHelpPrepend
	// message string
	h.message = message
	// subcommands    []HelpSubcommand
	for _, sc := range sc.Subcommands {
		newHelpSubcommand := HelpSubcommand{
			ShortName:   sc.ShortName,
			LongName:    sc.Name,
			Description: sc.Description,
		}
		h.subcommands = append(h.subcommands, newHelpSubcommand)
	}
	// positionals    []HelpPositional
	for _, pos := range sc.PositionalFlags {
		newHelpPositional := HelpPositional{
			Name:        pos.Name,
			Position:    pos.Position,
			Description: pos.Description,
			Required:    pos.Required,
		}
		h.positionals = append(h.positionals, newHelpPositional)
	}
	// flags          []HelpFlag
	for _, f := range sc.StringFlags {
		newHelpFlag := HelpFlag{
			ShortName:   f.ShortName,
			LongName:    f.LongName,
			Description: f.Description,
		}
		h.stringFlags = append(h.stringFlags, newHelpFlag)
	}
	for _, f := range sc.IntFlags {
		newHelpFlag := HelpFlag{
			ShortName:   f.ShortName,
			LongName:    f.LongName,
			Description: f.Description,
		}
		h.intFlags = append(h.intFlags, newHelpFlag)
	}
	for _, f := range sc.BoolFlags {
		newHelpFlag := HelpFlag{
			ShortName:   f.ShortName,
			LongName:    f.LongName,
			Description: f.Description,
		}
		h.boolFlags = append(h.boolFlags, newHelpFlag)
	}
}
