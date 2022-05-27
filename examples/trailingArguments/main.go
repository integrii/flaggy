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

	// add a global bool flag for fun
	flaggy.Bool(&someBool, "y", "yes", "A sample boolean flag")
	flaggy.String(&someString, "s", "string", "A sample string flag")
	flaggy.Int(&someInt, "i", "int", "A sample int flag")

	flaggy.ShowHelpOnUnexpectedDisable()

	// Parse the subcommand and all flags
	flaggy.Parse()

	fmt.Println(flaggy.TrailingArguments)
}
