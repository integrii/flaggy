package flaggy_test

import (
	"fmt"
	"os"

	"github.com/integrii/flaggy"
)

// ExampleSubcommand_AddPositionalValue adds two levels of subcommands with a
// positional value on the second level one
func ExampleSubcommand_AddPositionalValue() {

	// Simulate some input from the CLI.  Don't do this in your program.
	flaggy.ResetParser()
	os.Args = []string{"binaryName", "subcommandA", "subcommandB", "subcommandBPositionalValue"}

	// Imagine the following program usage:
	//
	// ./binaryName subcommandA subcommandB subcommandBPositional
	//

	var subcommandBPositional string

	// create a subcommand
	subcommandA := flaggy.NewSubcommand("subcommandA")
	// add the subcommand at relative position 1 within the default root parser
	flaggy.AttachSubcommand(subcommandA, 1)

	// create a second subcommand
	subcommandB := flaggy.NewSubcommand("subcommandB")
	// add the second subcommand to the first subcommand as a child at relative
	// position 1
	subcommandA.AttachSubcommand(subcommandB, 1)
	// add a positional to the second subcommand with a relative position of 1
	subcommandB.AddPositionalValue(&subcommandBPositional, "subcommandTestPositonalValue", 1, false, "A test positional input variable")

	// Parse the input arguments from the OS (os.Args) using the default parser
	flaggy.Parse()

	// see if our flag was set properly
	fmt.Println("Positional flag set to", subcommandBPositional)
	// Output: Positional flag set to subcommandBPositionalValue
}

// ExamplePositionalValue shows how to add positional variables at the
// global level.
func ExamplePositionalValue() {

	// Simulate some input from the CLI.  Don't do this in your program.
	flaggy.ResetParser()
	os.Args = []string{"binaryName", "positionalValue"}

	// Imagine the following program usage:
	//
	// ./binaryName positionalValue

	// add a bool flag at the global level
	var stringVar string
	flaggy.AddPositionalValue(&stringVar, "positionalVar", 1, false, "A test positional flag")

	// Parse the input arguments from the OS (os.Args)
	flaggy.Parse()

	// see if our flag was set properly
	if stringVar == "positionalValue" {
		fmt.Println("Flag set to", stringVar)
	}
	// Output: Flag set to positionalValue
}

// ExampleBoolFlag shows how to global bool flags in your program.
func ExampleBool() {

	// Simulate some input from the CLI.  Don't do these two lines in your program.
	flaggy.ResetParser()
	os.Args = []string{"binaryName", "-f"}

	// Imagine the following program usage:
	//
	// ./binaryName -f
	// or
	// ./binaryName --flag=true

	// add a bool flag at the global level
	var boolFlag bool
	flaggy.Bool(&boolFlag, "f", "flag", "A test bool flag")

	// Parse the input arguments from the OS (os.Args)
	flaggy.Parse()

	// see if our flag was set properly
	if boolFlag == true {
		fmt.Println("Flag set")
	}
	// Output: Flag set
}

// ExampleIntFlag shows how to global int flags in your program.
func ExampleInt() {

	// Simulate some input from the CLI.  Don't do these two lines in your program.
	flaggy.ResetParser()
	os.Args = []string{"binaryName", "-f", "5"}

	// Imagine the following program usage:
	//
	// ./binaryName -f 5
	// or
	// ./binaryName --flag=5

	// add a int flag at the global level
	var intFlag int
	flaggy.Int(&intFlag, "f", "flag", "A test int flag")

	// Parse the input arguments from the OS (os.Args)
	flaggy.Parse()

	// see if our flag was set properly
	if intFlag == 5 {
		fmt.Println("Flag set to:", intFlag)
	}
	// Output: Flag set to: 5
}

// Example shows how to add string flags in your program.
func Example() {

	// Simulate some input from the CLI.  Don't do this in your program.
	flaggy.ResetParser()
	os.Args = []string{"binaryName", "-f", "flagName"}

	// Imagine the following program usage:
	//
	// ./binaryName -f flagName
	// or
	// ./binaryName --flag=flagName

	// add a string flag at the global level
	var stringFlag string
	flaggy.String(&stringFlag, "f", "flag", "A test string flag")

	// Parse the input arguments from the OS (os.Args)
	flaggy.Parse()

	// see if our flag was set properly
	if stringFlag == "flagName" {
		fmt.Println("Flag set to:", stringFlag)
	}
	// Output: Flag set to: flagName
}

// ExampleSubcommand shows usage of subcommands in flaggy.
func ExampleSubcommand() {

	// Do not include the following two lines in your real program, it is for this
	// example only:
	flaggy.ResetParser()
	os.Args = []string{"programName", "-v", "VariableHere", "subcommandName", "subcommandPositional", "--", "trailingVar"}

	// Imagine the input to this program is as follows:
	//
	// ./programName subcommandName -v VariableHere subcommandPositional -- trailingVar
	//   or
	// ./programName subcommandName subcommandPositional --variable VariableHere -- trailingVar
	//   or
	// ./programName subcommandName --variable=VariableHere subcommandPositional -- trailingVar
	//   or even
	// ./programName subcommandName subcommandPositional -v=VariableHere -- trailingVar
	//

	// Create a new subcommand to attach flags and other subcommands to.  It must be attached
	// to something before being used.
	newSC := flaggy.NewSubcommand("subcommandName")

	// Attach a string variable to the subcommand
	var subcommandVariable string
	newSC.String(&subcommandVariable, "v", "variable", "A test variable.")

	var subcommandPositional string
	newSC.AddPositionalValue(&subcommandPositional, "testPositionalVar", 1, false, "A test positional variable to a subcommand.")

	// Attach the subcommand to the parser. This will panic if another
	// positional value or subcommand is already present at the depth supplied.
	// Later you can check if this command was used with a simple bool (newSC.Used).
	flaggy.AttachSubcommand(newSC, 1)

	// Parse the input arguments from the OS (os.Args)
	flaggy.Parse()

	// see if the subcommand was found during parsing:
	if newSC.Used {
		// Do subcommand operations here
		fmt.Println("Subcommand used")

		// check the input on your subcommand variable
		if subcommandVariable == "VariableHere" {
			fmt.Println("Subcommand variable set correctly")
		}

		// Print the subcommand positional value
		fmt.Println("Subcommand Positional:", subcommandPositional)

		// Print the first trailing argument
		fmt.Println("Trailing variable 1:", flaggy.TrailingArguments[0])
	}
	// Output:
	// Subcommand used
	// Subcommand variable set correctly
	// Subcommand Positional: subcommandPositional
	// Trailing variable 1: trailingVar
}
