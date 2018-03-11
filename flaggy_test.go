package flaggy_test

import (
	"testing"

	"github.com/integrii/flaggy"
)

func TestParsePositionalsA(t *testing.T) {
	flaggy.DebugMode = true
	inputLine := []string{"-t", "-i=3", "subcommand", "-n", "testN", "-j=testJ", "positionalA", "positionalB", "--testK=testK", "--", "trailingA", "trailingB"}

	var boolT bool
	var intT int
	var testN string
	var testJ string
	var testK string
	var positionalA string
	var positionalB string

	// make a new parser
	parser := flaggy.NewParser("testParser")

	// add a bool flag to the parser
	parser.AddBoolFlag(&boolT, "t", "", "test flag for bool arg")
	// add an int flag to the parser
	parser.AddIntFlag(&intT, "i", "", "test flag for int arg")

	// create a subcommand
	subCommand := flaggy.NewSubcommand("testSubcommand", 1)
	err := parser.AddSubcommand(subCommand)
	if err != nil {
		t.Fatal(err)
	}

	// add flags to subcommand
	subCommand.AddStringFlag(&testN, "n", "testN", "test flag for value with space arg")
	subCommand.AddStringFlag(&testJ, "j", "testJ", "test flag for value with equals arg")
	subCommand.AddStringFlag(&testK, "k", "testK", "test full length flag with attached arg")

	// add positionals to subcommand
	err = subCommand.AddPositionalValue(1, &positionalA, "PositionalA", "PositionalA test value")
	if err != nil {
		t.Fatal(err)
	}
	err = subCommand.AddPositionalValue(2, &positionalB, "PositionalB", "PositionalB test value")
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
