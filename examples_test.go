package flaggy_test

import (
	"fmt"
	"log"
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
	var err error

	// create a subcommand
	subcommandA := flaggy.NewSubcommand("subcommandA")
	// add the subcommand at relative positon 1 within the default root parser
	err = flaggy.AddSubcommand(subcommandA, 1)
	if err != nil {
		log.Fatal(err)
	}

	// create a second subcommand
	subcommandB := flaggy.NewSubcommand("subcommandB")
	// add the second subcommand to the first subcommand as a child at relative
	// position 1
	err = subcommandA.AddSubcommand(subcommandB, 1)
	if err != nil {
		log.Fatal(err)
	}
	// add a positional to the second subcommand with a relative position of 1
	err = subcommandB.AddPositionalValue(&subcommandBPositional, "subcommandTestPositonalValue", 1, false, "A test positional input variable")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the input arguments from the OS (os.Args) using the default parser
	err = flaggy.Parse()
	if err != nil {
		log.Fatal(err)
	}

	// see if our flag was set properly
	fmt.Println("Positional flag set to", subcommandBPositional)
	// Output: Positional flag set to subcommandBPositionalValue
}

// ExampleAddPositionalValue shows how to add positional vairables at the
// global level.
func ExampleAddPositionalValue() {

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

// ExampleAddBoolFlag shows how to global bool flags in your program.
func ExampleAddBoolFlag() {

	// Simulate some input from the CLI.  Don't do this in your program.
	flaggy.ResetParser()
	os.Args = []string{"binaryName", "-f"}

	// Imagine the following program usage:
	//
	// ./binaryName -f
	// or
	// ./binaryName --flag=true

	// add a bool flag at the global level
	var boolFlag bool
	flaggy.AddBoolFlag(&boolFlag, "f", "flag", "A test bool flag")

	// Parse the input arguments from the OS (os.Args)
	flaggy.Parse()

	// see if our flag was set properly
	if boolFlag == true {
		fmt.Println("Flag set")
	}
	// Output: Flag set
}

// ExampleAddIntFlag shows how to global int flags in your program.
func ExampleAddIntFlag() {

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
	flaggy.AddIntFlag(&intFlag, "f", "flag", "A test int flag")

	// Parse the input arguments from the OS (os.Args)
	flaggy.Parse()

	// see if our flag was set properly
	if intFlag == 5 {
		fmt.Println("Flag set to:", intFlag)
	}
	// Output: Flag set to: 5
}

// ExampleAddStringFlag shows how to global string flags in your program.
func ExampleAddStringFlag() {

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
	flaggy.AddStringFlag(&stringFlag, "f", "flag", "A test string flag")

	// Parse the input arguments from the OS (os.Args)
	flaggy.Parse()

	// see if our flag was set properly
	if stringFlag == "flagName" {
		fmt.Println("Flag set to:", stringFlag)
	}
	// Output: Flag set to: flagName
}

// Example shows some basic usage of flaggy.
func Example() {

	// Do not include the following line in your real program, it is for this
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

	// Create a new subcommand at the first position to what it is attached to.
	// The depth is relative to the thing this subcommand is attached to.
	newSC := flaggy.NewSubcommand("subcommandName")

	// Attach a string variable to the subcommand
	var subcommandVariable string
	newSC.AddStringFlag(&subcommandVariable, "v", "variable", "A test variable.")

	var subcommandPositional string
	newSC.AddPositionalValue(&subcommandPositional, "testPositionalVar", 1, false, "A test positional variable to a subcommand.")

	// Attach the subcommand to the parser. This will error if another
	// positional value or subcommand is already present at the depth supplied.
	// Later you can check if this command was used with a simple bool.
	err := flaggy.AddSubcommand(newSC, 1)
	if err != nil {
		log.Fatalln(err)
	}

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
