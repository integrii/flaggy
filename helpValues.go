package flaggy

// help represents the values needed to render a help page
type help struct {
	subcommands    []helpSubcommand
	positionals    []helpPositional
	stringFlags    []helpFlag
	intFlags       []helpFlag
	boolFlags      []helpFlag
	prependMessage string
	appendMessage  string
	message        string
}

// helpSubcommand is used to template subcommand help output
type helpSubcommand struct {
	ShortName   string
	LongName    string
	Description string
}

// helpPositional is used to template positional help output
type helpPositional struct {
	Name        string
	Description string
	Required    bool
	Position    int
}

// helpFlag is used to template string flag help output
type helpFlag struct {
	ShortName   string
	LongName    string
	Description string
	FlagType    string
}

const helpFlagTypeBool = "bool"
const helpFlagTypeString = "string"
const helpFlagTypeInt = "int"

// ExtractValuesFromSubcommand extracts help template values from a subcommand
func (h *help) ExtractValuesFromSubcommand(sc *Subcommand, message string) {
	// extract help values
	// prependMessage string
	h.prependMessage = sc.AdditionalHelpPrepend
	// appendMessage  string
	h.appendMessage = sc.AdditionalHelpPrepend
	// message string
	h.message = message
	// subcommands    []helpSubcommand
	for _, sc := range sc.Subcommands {
		newHelpSubcommand := helpSubcommand{
			ShortName:   sc.ShortName,
			LongName:    sc.Name,
			Description: sc.Description,
		}
		h.subcommands = append(h.subcommands, newHelpSubcommand)
	}
	// positionals    []helpPositional
	for _, pos := range sc.PositionalFlags {
		newHelpPositional := helpPositional{
			Name:        pos.Name,
			Position:    pos.Position,
			Description: pos.Description,
			Required:    pos.Required,
		}
		h.positionals = append(h.positionals, newHelpPositional)
	}
	// flags          []helpFlag
	for _, f := range sc.StringFlags {
		newHelpFlag := helpFlag{
			ShortName:   f.ShortName,
			LongName:    f.LongName,
			Description: f.Description,
			FlagType:    helpFlagTypeString,
		}
		h.stringFlags = append(h.stringFlags, newHelpFlag)
	}
	for _, f := range sc.IntFlags {
		newHelpFlag := helpFlag{
			ShortName:   f.ShortName,
			LongName:    f.LongName,
			Description: f.Description,
			FlagType:    helpFlagTypeInt,
		}
		h.intFlags = append(h.intFlags, newHelpFlag)
	}
	for _, f := range sc.BoolFlags {
		newHelpFlag := helpFlag{
			ShortName:   f.ShortName,
			LongName:    f.LongName,
			Description: f.Description,
			FlagType:    helpFlagTypeBool,
		}
		h.boolFlags = append(h.boolFlags, newHelpFlag)
	}
}
