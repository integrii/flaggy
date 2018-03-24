package flaggy_test

import (
	"testing"

	"github.com/integrii/flaggy"
)

// TestComplexNesting tests various levels of nested subcommands and
// positional values intermixed with eachother.
func TestComplexNesting(t *testing.T) {
	var testA string
	var testB string
	var testC string
	var testD string
	var testE string

	scA := flaggy.NewSubcommand("scA")
	scB := flaggy.NewSubcommand("scB")
	scC := flaggy.NewSubcommand("scC")
	scD := flaggy.NewSubcommand("scD")

	scA.AddPositionalValue(&testA, "testA", 1, false, "")
	scA.AddPositionalValue(&testB, "testB", 2, false, "")
	scA.AddPositionalValue(&testC, "testC", 3, false, "")
	scA.AddSubcommand(scB, 4)
	flaggy.AddSubcommand(scA, 1)

	scB.AddPositionalValue(&testD, "testD", 1, false, "")
	scB.AddSubcommand(scC, 2)

	scC.AddSubcommand(scD, 1)

	scD.AddPositionalValue(&testE, "testE", 1, true, "")

	flaggy.ParseArgs([]string{"scA", "A", "B", "C", "scB", "D", "scC", "scD", "E"})

	if !scA.Used {
		t.FailNow()
	}
	if !scB.Used {
		t.FailNow()
	}
	if !scC.Used {
		t.FailNow()
	}
	if !scD.Used {
		t.FailNow()
	}
	if testA != "A" {
		t.FailNow()
	}
	if testB != "B" {
		t.FailNow()
	}
	if testC != "C" {
		t.FailNow()
	}
	if testD != "D" {
		t.FailNow()
	}
	if testE != "E" {
		t.FailNow()
	}

}

func TestParsePositionalsA(t *testing.T) {
	inputLine := []string{"-t", "-i=3", "subcommand", "-n", "testN", "-j=testJ", "positionalA", "positionalB", "--testK=testK", "--", "trailingA", "trailingB"}

	var boolT bool
	var intT int
	var testN string
	var testJ string
	var testK string
	var positionalA string
	var positionalB string
	var err error

	// make a new parser
	parser := flaggy.NewParser("testParser")

	// add a bool flag to the parser
	err = parser.AddBoolFlag(&boolT, "t", "", "test flag for bool arg")
	if err != nil {
		t.Fatal(err)
	}
	// add an int flag to the parser
	err = parser.AddIntFlag(&intT, "i", "", "test flag for int arg")
	if err != nil {
		t.Fatal(err)
	}

	// create a subcommand
	subCommand := flaggy.NewSubcommand("subcommand")
	err = parser.AddSubcommand(subCommand, 1)
	if err != nil {
		t.Fatal(err)
	}

	// add flags to subcommand
	err = subCommand.AddStringFlag(&testN, "n", "testN", "test flag for value with space arg")
	if err != nil {
		t.Fatal(err)
	}
	err = subCommand.AddStringFlag(&testJ, "j", "testJ", "test flag for value with equals arg")
	if err != nil {
		t.Fatal(err)
	}
	err = subCommand.AddStringFlag(&testK, "k", "testK", "test full length flag with attached arg")
	if err != nil {
		t.Fatal(err)
	}

	// add positionals to subcommand
	err = subCommand.AddPositionalValue(&positionalA, "PositionalA", 1, false, "PositionalA test value")
	if err != nil {
		t.Fatal(err)
	}
	err = subCommand.AddPositionalValue(&positionalB, "PositionalB", 2, false, "PositionalB test value")
	if err != nil {
		t.Fatal(err)
	}

	// parse input
	err = parser.ParseArgs(inputLine)
	if err != nil {
		t.Fatal(err)
	}

	// check the results
	if intT != 3 {
		t.Fatal("Global int flag -i was incorrect:", intT)
	}
	if boolT != true {
		t.Fatal("Global bool flag -t was incorrect:", boolT)
	}
	if testK != "testK" {
		t.Fatal("Subcommand flag testK was incorrect:", testK)
	}
	if testN != "testN" {
		t.Fatal("Subcommand flag testN was incorrect:", testN)
	}
	if testJ != "testJ" {
		t.Fatal("Subcommand flag testJ was incorrect:", testJ)
	}
	if positionalA != "positionalA" {
		t.Fatal("Positional A was incorrect:", positionalA)
	}
	if positionalB != "positionalB" {
		t.Fatal("Positional B was incorrect:", positionalB)
	}
	if len(parser.TrailingArguments) < 2 {
		t.Fatal("Incorrect number of trailing arguments.  Got", len(parser.TrailingArguments))
	}
	if parser.TrailingArguments[0] != "trailingA" {
		t.Fatal("Trailing argumentA was incorrect:", parser.TrailingArguments[0])
	}
	if parser.TrailingArguments[1] != "trailingB" {
		t.Fatal("Trailing argumentB was incorrect:", parser.TrailingArguments[1])
	}

}
