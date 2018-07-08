package main

import (
	"fmt"

	"github.com/integrii/flaggy"
)

// Declare variables and their defaults
var positionalValue = "defaultString"
var intFlagT = 3
var boolFlagB bool

func main() {

	// set a description, name, and version for our parser
	p := flaggy.NewParser("myAppName")
	p.Description = "This parser just shows you how to make a parser."
	p.Version = "1.3.5"
	// display some before and after text for all help outputs
	p.AdditionalHelpPrepend = "I hope you like this program!"
	p.AdditionalHelpAppend = "This command has no warranty."

	// add a positional value at position 1
	p.AddPositionalValue(&positionalValue, "testPositional", 1, true, "This is a test positional value that is required")

	// create a subcommand at position 2
	// you don't have to finish the subcommand before adding it to the parser
	subCmd := flaggy.NewSubcommand("subCmd")
	subCmd.Description = "Description of subcommand"
	p.AttachSubcommand(subCmd, 2)

	// add a flag to the subcomand
	subCmd.Int(&intFlagT, "i", "testInt", "This is a test int flag")

	// add a bool flag to the root command
	p.Bool(&boolFlagB, "b", "boolTest", "This is a test boolean flag")

	p.Parse()

	fmt.Println(positionalValue, intFlagT, boolFlagB)

	// Imagine the following command line:
	// ./customParser positionalHere subCmd -i 33 -b
	// It would produce:
	// positionalHere 33 true
}
