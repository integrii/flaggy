package flaggy

// PositionalValue represents a value which is determined by its position
// relative to where a subcommand was detected.
type PositionalValue struct {
	Name          string // used in documentation only
	Description   string
	AssignmentVar *string // the var that will get this variable
	Position      int     // the position, not including switches, of this variable
}
