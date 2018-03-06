package flaggy

import (
	"fmt"
	"os"
)

// PositionalValue represents a value which is determined by its position
// relative to where a subcommand was detected.
type PositionalValue struct {
	Name          string // used in documentation only
	Description   string
	AssignmentVar *string // the var that will get this variable
	Position      int     // the position, not including switches, of this variable
}

func NewPositionalValue(relativeDepth int) {
	if relativeDepth < 1 {
		fmt.Println("Flaggy: Position of positional arguments must never be below 1")
		os.Exit(2)
	}
}
