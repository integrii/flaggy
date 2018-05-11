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

	//  a flag to the subcommand
	subcommandExample.String(&stringFlagF, "t", "testFlag", "A test string flag")
	nestedSubcommand.Int(&intFlagT, "f", "flag", "A test int flag")

	// add a global bool flag for fun
	flaggy.Bool(&boolFlagB, "y", "yes", "A sample boolean flag")

	//  the nested subcommand to the parent subcommand at position 1
	subcommandExample.AttachSubcommand(nestedSubcommand, 1)

	//  the base subcommand to the parser at position 1
	flaggy.AttachSubcommand(subcommandExample, 1)

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
