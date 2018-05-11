package main

import "github.com/integrii/flaggy"

func main() {
	// Declare variables and their defaults
	var stringFlag = "defaultValue"

	// Create the subcommand
	subcommand := flaggy.NewSubcommand("subcommandExample")

	// Add a flag to the subcommand
	subcommand.String(&stringFlag, "f", "flag", "A test string flag")

	//  the subcommand to the parser at position 1
	flaggy.AttachSubcommand(subcommand, 1)

	// Parse the subcommand and all flags
	flaggy.Parse()

	// Use the flag
	print(stringFlag)
}
