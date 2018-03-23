package main

import "github.com/integrii/flaggy"

func main() {

	// Declare variables and their defaults
	var stringFlagF = "defaultValueF"
	var intFlagT = 3
	var boolFlagB bool

	// Create the subcommand
	subcommandExample := flaggy.NewSubcommand("subcommandExample")
	nestedSubcommand := flaggy.NewSubcommand("nestedSubcommand")

	// Add a flag to the subcommand
	subcommandExample.AddStringFlag(&stringFlagF, "t", "testFlag", "A test string flag")
	nestedSubcommand.AddIntFlag(&intFlagT, "f", "flag", "A test int flag")

	// add a global bool flag for fun
	flaggy.AddBoolFlag(&boolFlagB, "y", "yes", "A sample boolean flag")

	// Add the nested subcommand to the parent subcommand at position 1
	subcommandExample.AddSubcommand(nestedSubcommand, 1)

	// Add the base subcommand to the parser at position 1
	flaggy.AddSubcommand(subcommandExample, 1)

	// Parse the subcommand and all flags
	flaggy.Parse()

	// Use the flags and trailing arguments
	print(stringFlagF)
	print(intFlagT)

	// we can check if a subcommand was used easily
	if nestedSubcommand.Used {
		print(boolFlagB)
	}
	print(flaggy.TrailingArguments[0:])
}
