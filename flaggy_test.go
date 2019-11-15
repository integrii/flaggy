package flaggy_test

import (
	"os"
	"testing"

	"github.com/integrii/flaggy"
)

// TestTrailingArguments tests trailing argument parsing
func TestTrailingArguments(t *testing.T) {
	flaggy.ResetParser()
	args := []string{"./flaggy.text", "--", "one", "two"}
	os.Args = args
	flaggy.Parse()
	if len(flaggy.TrailingArguments) != 2 {
		t.Fatal("incorrect argument count parsed.  Got", len(flaggy.TrailingArguments), "but expected", 2)
	}

	if flaggy.TrailingArguments[0] != "one" {
		t.Fatal("incorrect argument parsed.  Got", flaggy.TrailingArguments[0], "but expected one")
	}

	if flaggy.TrailingArguments[1] != "two" {
		t.Fatal("incorrect argument parsed.  Got", flaggy.TrailingArguments[1], "but expected two")
	}

}

// TestComplexNesting tests various levels of nested subcommands and
// positional values intermixed with eachother.
func TestComplexNesting(t *testing.T) {

	flaggy.DebugMode = true
	defer debugOff()

	flaggy.ResetParser()

	var testA string
	var testB string
	var testC string
	var testD string
	var testE string
	var testF bool

	scA := flaggy.NewSubcommand("scA")
	scB := flaggy.NewSubcommand("scB")
	scC := flaggy.NewSubcommand("scC")
	scD := flaggy.NewSubcommand("scD")

	flaggy.Bool(&testF, "f", "testF", "")

	flaggy.AttachSubcommand(scA, 1)

	scA.AddPositionalValue(&testA, "testA", 1, false, "")
	scA.AddPositionalValue(&testB, "testB", 2, false, "")
	scA.AddPositionalValue(&testC, "testC", 3, false, "")
	scA.AttachSubcommand(scB, 4)

	scB.AddPositionalValue(&testD, "testD", 1, false, "")
	scB.AttachSubcommand(scC, 2)

	scC.AttachSubcommand(scD, 1)

	scD.AddPositionalValue(&testE, "testE", 1, true, "")

	args := []string{"scA", "-f", "A", "B", "C", "scB", "D", "scC", "scD", "E"}
	t.Log(args)
	flaggy.ParseArgs(args)

	if !testF {
		t.Log("testF", testF)
		t.FailNow()
	}
	if !scA.Used {
		t.Log("sca", scA.Name)
		t.FailNow()
	}
	if !scB.Used {
		t.Log("scb", scB.Name)
		t.FailNow()
	}
	if !scC.Used {
		t.Log("scc", scC.Name)
		t.FailNow()
	}
	if !scD.Used {
		t.Log("scd", scD.Name)
		t.FailNow()
	}
	if testA != "A" {
		t.Log("testA", testA)
		t.FailNow()
	}
	if testB != "B" {
		t.Log("testB", testB)
		t.FailNow()
	}
	if testC != "C" {
		t.Log("testC", testC)
		t.FailNow()
	}
	if testD != "D" {
		t.Log("testD", testD)
		t.FailNow()
	}
	if testE != "E" {
		t.Log("testE", testE)
		t.FailNow()
	}

}

func TestParsePositionalsA(t *testing.T) {
	inputLine := []string{"-t", "-i=3", "subcommand", "-n", "testN", "-j=testJ", "positionalA", "positionalB", "--testK=testK", "--", "trailingA", "trailingB"}

	flaggy.DebugMode = true

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
	parser.Bool(&boolT, "t", "", "test flag for bool arg")
	// add an int flag to the parser
	parser.Int(&intT, "i", "", "test flag for int arg")

	// create a subcommand
	subCommand := flaggy.NewSubcommand("subcommand")
	parser.AttachSubcommand(subCommand, 1)

	// add flags to subcommand
	subCommand.String(&testN, "n", "testN", "test flag for value with space arg")
	subCommand.String(&testJ, "j", "testJ", "test flag for value with equals arg")
	subCommand.String(&testK, "k", "testK", "test full length flag with attached arg")

	// add positionals to subcommand
	subCommand.AddPositionalValue(&positionalA, "PositionalA", 1, false, "PositionalA test value")
	subCommand.AddPositionalValue(&positionalB, "PositionalB", 2, false, "PositionalB test value")

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
	if testN != "testN" {
		t.Fatal("Subcommand flag testN was incorrect:", testN)
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
