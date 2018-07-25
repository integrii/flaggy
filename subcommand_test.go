package flaggy_test

import (
	"os"
	"testing"

	"github.com/integrii/flaggy"
)

// TestDoublePositional tests errors when two positionals are
// specified at the same time
func TestDoublePositional(t *testing.T) {
	t.Skip("Skipped.  If this test runs, it exists the whole program.")
	// flaggy.DebugMode = true
	defer debugOff()
	var posTest string
	flaggy.ResetParser()
	flaggy.AddPositionalValue(&posTest, "posTest", 1, true, "First test positional")
	flaggy.AddPositionalValue(&posTest, "posTest2", 1, true, "Second test positional")
}

// TestRequiredPositional tests required positionals
func TestRequiredPositional(t *testing.T) {
	t.Skip("Skipped.  If this test runs, it exists the whole program.")
	// flaggy.DebugMode = true
	defer debugOff()
	var posTest string
	flaggy.AddPositionalValue(&posTest, "posTest", 1, true, "First test positional")
	flaggy.Parse()
}

// TestTypoSubcommand tests what happens when an invalid subcommand is passed
func TestTypoSubcommand(t *testing.T) {
	t.Skip("Skipped.  If this test runs, it exists the whole program.")
	p := flaggy.NewParser("TestTypoSubcommand")
	p.ShowHelpOnUnexpected = true
	args := []string{"unexpectedArg"}
	newSCA := flaggy.NewSubcommand("TestTypoSubcommandA")
	newSCB := flaggy.NewSubcommand("TestTypoSubcommandB")
	p.AttachSubcommand(newSCA, 1)
	p.AttachSubcommand(newSCB, 1)
	p.ParseArgs(args)
}

// TestIgnoreUnexpected tests what happens when an invalid subcommand is passed but should be ignored
func TestIgnoreUnexpected(t *testing.T) {
	p := flaggy.NewParser("TestTypoSubcommand")
	p.ShowHelpOnUnexpected = false
	args := []string{"unexpectedArg"}
	newSCA := flaggy.NewSubcommand("TestTypoSubcommandA")
	p.AttachSubcommand(newSCA, 1)
	p.ParseArgs(args)
}

// TestSubcommandHelp tests displaying of help on unspecified commands
func TestSubcommandHelp(t *testing.T) {
	t.Skip("Skipped.  If this test runs, it exists the whole program.")
	p := flaggy.NewParser("TestSubcommandHelp")
	p.ShowHelpOnUnexpected = true
	args := []string{"unexpectedArg"}
	p.ParseArgs(args)
}

func TestHelpWithHFlagA(t *testing.T) {
	t.Skip("Skipped.  If this test runs, it exists the whole program.")
	p := flaggy.NewParser("TestHelpWithHFlag")
	p.ShowHelpWithHFlag = true
	args := []string{"-h"}
	p.ParseArgs(args)
}

func TestHelpWithHFlagB(t *testing.T) {
	t.Skip("Skipped.  If this test runs, it exists the whole program.")
	p := flaggy.NewParser("TestHelpWithHFlag")
	p.ShowHelpWithHFlag = true
	args := []string{"--help"}
	p.ParseArgs(args)
}

func TestVersionWithVFlagB(t *testing.T) {
	t.Skip("Skipped.  If this test runs, it exists the whole program.")
	p := flaggy.NewParser("TestSubcommandVersion")
	p.ShowVersionWithVersionFlag = true
	p.Version = "TestVersionWithVFlagB 0.0.0a"
	args := []string{"--version"}
	p.ParseArgs(args)
}

// TestSubcommandParse tests paring of a single subcommand
func TestSubcommandParse(t *testing.T) {

	var positionA string

	// create the argument parser
	p := flaggy.NewParser("TestSubcommandParse")

	// create a subcommand
	newSC := flaggy.NewSubcommand("testSubcommand")

	// add the subcommand into the parser
	err := p.AttachSubcommand(newSC, 1)
	if err != nil {
		t.Fatal("Error adding subcommand", err)
	}

	// add a positional arg onto the subcommand at relative position 1
	err = newSC.AddPositionalValue(&positionA, "positionalA", 1, false, "This is a test positional value")
	if err != nil {
		t.Fatal("Error adding positional value", err)
	}

	// override os args and parse them
	os.Args = []string{"binaryName", "testSubcommand", "testPositional"}
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

	// create the argument parser
	p := flaggy.NewParser("TestBadSubcommand")

	// create a subcommand
	newSC := flaggy.NewSubcommand("testSubcommand")
	err := p.AttachSubcommand(newSC, 1)
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

	// create the argument parser
	p := flaggy.NewParser("TestBadPositional")

	// create a subcommand
	// add a positional arg into the subcommand
	var positionA string
	var err error
	err = p.AddPositionalValue(&positionA, "positionalA", 1, false, "This is a test positional value")
	if err != nil {
		t.Fatal(err)
	}

	//  test what happens if you add a bad subcommand
	os.Args = []string{"test", "badPositional"}
	err = p.Parse()
	if err != nil {
		t.Fatal("Threw an error when bad positional was passed, but shouldn't have")
	}
}

// TestNakedBoolFlag tests a naked boolean flag, which mean it has no
// specified value beyond the flag being present.
func TestNakedBool(t *testing.T) {
	flaggy.ResetParser()
	os.Args = []string{"testBinary", "-t"}

	// add a bool var
	var boolVar bool
	flaggy.Bool(&boolVar, "t", "boolVar", "A boolean flag for testing")
	flaggy.Parse()
	if !boolVar {
		t.Fatal("Boolean naked val not set to true")
	}
}

// debugOff makes defers easier
func debugOff() {
	// flaggy.DebugMode = false
}

// BenchmarkSubcommandParse benchmarks the creation and parsing of
// a basic subcommand
func BenchmarkSubcommandParse(b *testing.B) {

	// catch errors that may occur
	defer func(b *testing.B) {
		err := recover()
		if err != nil {
			b.Fatal("Benchmark had error:", err)
		}
	}(b)

	for i := 0; i < b.N; i++ {

		var positionA string

		// create the argument parser
		p := flaggy.NewParser("TestSubcommandParse")

		// create a subcommand
		newSC := flaggy.NewSubcommand("testSubcommand")

		// add the subcommand into the parser
		err := p.AttachSubcommand(newSC, 1)
		if err != nil {
			b.Fatal("Error adding subcommand", err)
		}

		// add a positional arg onto the subcommand at relative position 1
		err = newSC.AddPositionalValue(&positionA, "positionalA", 1, false, "This is a test positional value")
		if err != nil {
			b.Fatal("Error adding positional value", err)
		}

		// override os args and parse them
		os.Args = []string{"binaryName", "testSubcommand", "testPositional"}
		err = p.Parse()
		if err != nil {
			b.Fatal("Error parsing args: " + err.Error())
		}
	}

}
