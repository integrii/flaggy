package flaggy_test

import (
	"os"
	"testing"

	"github.com/integrii/flaggy"
)

// TestSubcommandParse tests paring of a single subcommand
func TestSubcommandParse(t *testing.T) {
	flaggy.DebugMode = true
	defer debugOff()
	var positionA string

	// create the argument parser
	p := flaggy.NewParser("TestSubcommandParse")

	// create a subcommand
	newSC := flaggy.NewSubcommand("testSubcommand")
	// add a positional arg into the subcommand
	err := newSC.AddPositionalValue(&positionA, "positionalA", 1, "This is a test positional value")
	if err != nil {
		t.Fatal("Error adding positional value", err)
	}

	// add the subcommand into the parser
	err = p.AddSubcommand(newSC, 1)
	if err != nil {
		t.Fatal("Error adding subcommand", err)
	}

	// override os args and parse them
	os.Args = []string{"testSubcommand", "testPositional"}
	err = p.Parse()
	if err != nil {
		t.Fatal("Error parsing args: " + err.Error())
	}

	// ensure subcommand and positional used
	if !newSC.Used {
		t.Fatal("Subcommand was not used, but it was expected to be")
	}
	if positionA != "testPositional" {
		t.Fatal("Positional argument was not set to testPositional, was:", positionA)
	}
}

func TestBadSubcommand(t *testing.T) {
	flaggy.DebugMode = true
	defer debugOff()

	// create the argument parser
	p := flaggy.NewParser("TestBadSubcommand")

	// create a subcommand
	newSC := flaggy.NewSubcommand("testSubcommand")
	err := p.AddSubcommand(newSC, 1)
	if err != nil {
		t.Fatal("Error adding subcommand", err)
	}

	//  test what happens if you add a bad subcommand
	os.Args = []string{"test"}
	err = p.Parse()
	if err != nil {
		t.Fatal("Threw an error when bad subcommand positional was passed, but should not have")
	}
}

func TestBadPositional(t *testing.T) {
	flaggy.DebugMode = true
	defer debugOff()

	// create the argument parser
	p := flaggy.NewParser("TestBadPositional")

	// create a subcommand
	// add a positional arg into the subcommand
	var positionA string
	p.AddPositionalValue(&positionA, "positionalA", 1, "This is a test positional value")

	//  test what happens if you add a bad subcommand
	os.Args = []string{"test", "badPositional"}
	err := p.Parse()
	if err != nil {
		t.Fatal("Threw an error when bad positional was passed, but shouldn't have")
	}
}

// TestNakedBoolFlag tests naked boolean flags
func TestNakedBoolFlag(t *testing.T) {
	flaggy.DebugMode = true
	defer debugOff()
	os.Args = []string{"testBinary", "-t"}

	// add a bool var
	var boolVar bool
	flaggy.AddBoolFlag(&boolVar, "t", "boolVar", "A boolean flag for testing")

	err := flaggy.Parse()
	if err != nil {
		t.Fatal(err)
	}

	if !boolVar {
		t.Fatal("Boolean naked val not set to true")
	}
}

// debugOff makes defers easier
func debugOff() {
	flaggy.DebugMode = false
}
