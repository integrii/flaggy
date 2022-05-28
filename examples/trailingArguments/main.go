package main

import (
	"fmt"

	"github.com/integrii/flaggy"
)

func main() {

	// Declare variables and their defaults
	var someString = ""
	var someInt = 3
	var someBool bool
	var positionalValue string

	// add a global bool flag for fun
	flaggy.Bool(&someBool, "y", "yes", "A sample boolean flag")
	flaggy.String(&someString, "s", "string", "A sample string flag")
	flaggy.Int(&someInt, "i", "int", "A sample int flag")

	// this positional value will be parsed specifically before all trailing
	// arguments are parsed
	flaggy.AddPositionalValue(&positionalValue, "testPositional", 1, false, "a test positional")

	flaggy.DebugMode = false
	flaggy.ShowHelpOnUnexpectedDisable()

	// Parse the subcommand and all flags
	flaggy.Parse()

	// here you will see all arguments passsed after the first positional 'testPositional' string is parsed
	fmt.Println(flaggy.TrailingArguments)
	// Input:
	// ./trailingArguments one two three
	// Output:
	// [two three]
}
