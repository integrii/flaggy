package main

import (
	"github.com/integrii/flaggy"
)

var version = "demo"

var (
	cmdA *flaggy.Subcommand
	cmdB *flaggy.Subcommand
)

// Declare variables and their defaults
var (
	stringFlagA = "defaultValueA"
	stringFlagB = "defaultValueB"
)

func main() {
	// Set your program's name and description.  These appear in help output.
	flaggy.SetName("subcommandhelpdemo")
	flaggy.SetDescription("tool for demonstrating subcommand help")

	// You can disable various things by changing bools on the default parser
	// (or your own parser if you have created one).
	flaggy.DefaultParser.ShowHelpOnUnexpected = true

	// You can set a help prepend or append on the default parser.
	flaggy.DefaultParser.AdditionalHelpPrepend = "Extra (global) before text"

	// Add a flag to the root of flaggy
	flaggy.String(&stringFlagA, "a", "flagA", "A test string flag (A)")

	cmdA = flaggy.NewSubcommand("cmdA")
	cmdA.ShortName = "A"
	cmdB = flaggy.NewSubcommand("cmdB")
	cmdB.ShortName = "B"
	cmdB.String(&stringFlagB, "b", "flagB", "A test string flag (B)")
	cmdB.AdditionalHelpPrepend = "extra help text for b."

	flaggy.AttachSubcommand(cmdA, 1)
	flaggy.AttachSubcommand(cmdB, 1)

	// Set the version and parse all inputs into variables.
	flaggy.SetVersion(version)
	flaggy.Parse()

	switch {
	case cmdA.Used:
		cmdA.ShowHelp()
	case cmdB.Used:
		cmdB.ShowHelp()
	default:
		flaggy.ShowHelp("Master-Helptext.")
	}

}
