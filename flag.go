package flaggy

// flag holds the base methods for all flag types
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
