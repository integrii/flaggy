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
	newSC := flaggy.NewSubcommand("testSubcommand", 1)
	// add a positional arg into the subcommand
	err := newSC.AddPositionalValue(1, &positionA, "positionalA", "This is a test positional value")
	if err != nil {
		t.Fatal("Error adding positional value", err)
	}

	// add the subcommand into the parser
	err = p.AddSubcommand(newSC)
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
	newSC := flaggy.NewSubcommand("testSubcommand", 1)
	err := p.AddSubcommand(newSC)
	if err != nil {
		t.Fatal("Error adding subcommand", err)
	}

	//  test what happens if you add a bad subcommand
	os.Args = []string{"test"}
	err = p.Parse()
	if err == nil {
		t.Fatal("Threw no error when bad subcommand positional was passed")
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
	p.AddPositionalValue(1, &positionA, "positionalA", "This is a test positional value")

	//  test what happens if you add a bad subcommand
	os.Args = []string{"test", "badPositional"}
	err := p.Parse()
	if err == nil {
		t.Fatal("Threw no error when bad subcommand positional was passed")
	}
}

// debugOff makes defers easier
func debugOff() {
	flaggy.DebugMode = false
}
