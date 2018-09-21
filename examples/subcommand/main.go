package main

import "github.com/integrii/flaggy"

// Declare variables and their defaults
var stringFlagA = "defaultValueA"
var stringFlagB = "defaultValueB"

func main() {

	// Add a flag to the root of flaggy
	flaggy.String(&stringFlagA, "a", "flagA", "A test string flag (A)")

	// Create the subcommand
	subcommand := flaggy.NewSubcommand("subcommandExample")

	// Add a flag to the subcommand
	subcommand.String(&stringFlagB, "b", "flagB", "A test string flag (B)")

	// Add the subcommand to the parser at position 1
	flaggy.AttachSubcommand(subcommand, 1)

	// Parse the subcommand and all flags
	flaggy.Parse()

	// Use the flags
	println("A: " + stringFlagA)
	println("B: " + stringFlagB)
}
