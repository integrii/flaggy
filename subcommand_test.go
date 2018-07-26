package flaggy_test

import (
	"os"
	"testing"

	"github.com/integrii/flaggy"
)

func TestFlagExists(t *testing.T) {
	sc := flaggy.NewSubcommand("testFlagExists")
	e := sc.FlagExists("test")
	if e == true {
		t.Fatal("Flag exists on subcommand that should not")
	}
	var testA string

	sc.String(&testA, "", "test", "a test flag")
	e = sc.FlagExists("test")
	if e == false {
		t.Fatal("Flag does not exist on a subcommand that should")
	}

}

// func TestBoolStringSupplied(t *testing.T) {
// 	flaggy.ResetParser()
// 	flaggy.DebugMode = true
// 	defer debugOff()
// 	var boolA bool
// 	flaggy.Bool(&boolA, "b", "boolean", "test bool flag")
// 	os.Args = []string{"-b", "true"}
// 	flaggy.Parse()
// }

// TestDoublePositional tests errors when two positionals are
// specified at the same time
func TestDoublePositional(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Expected crash on double assignment")
		}
	}()
	// flaggy.DebugMode = true
	defer debugOff()
	var posTest string
	flaggy.ResetParser()
	flaggy.AddPositionalValue(&posTest, "posTest", 1, true, "First test positional")
	flaggy.AddPositionalValue(&posTest, "posTest2", 1, true, "Second test positional")
}

func TestNextArgDoesNotExist(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Expected crash when next arg not specifid")
		}
	}()
	flaggy.ResetParser()
	flaggy.PanicInsteadOfExit = true
	var test string
	flaggy.String(&test, "t", "test", "Description goes here")
	flaggy.ParseArgs([]string{"-t"})
}

func TestSubcommandHidden(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Expected crash instead of exit.  Subcommand id was wrong")
		}
	}()
	flaggy.ResetParser()
	sc := flaggy.NewSubcommand("")
	sc.Hidden = true
	sc.ShortName = "sc"
	flaggy.AttachSubcommand(sc, 1)
	flaggy.ParseArgs([]string{"x"})
}

// TestRequiredPositional tests required positionals
func TestRequiredPositional(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Expected crash on unused required positional")
		}
	}()
	// flaggy.DebugMode = true
	defer debugOff()
	var posTest string
	flaggy.AddPositionalValue(&posTest, "posTest", 1, true, "First test positional")
	flaggy.Parse()
}

// TestTypoSubcommand tests what happens when an invalid subcommand is passed
func TestTypoSubcommand(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Expected crash on subcommand typo")
		}
	}()
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
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Expected crash on subcommand help display")
		}
	}()
	p := flaggy.NewParser("TestSubcommandHelp")
	p.ShowHelpOnUnexpected = true
	args := []string{"unexpectedArg"}
	p.ParseArgs(args)
}

func TestHelpWithHFlagA(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Expected crash on help flag use")
		}
	}()
	p := flaggy.NewParser("TestHelpWithHFlag")
	p.ShowHelpWithHFlag = true
	args := []string{"-h"}
	p.ParseArgs(args)
}

func TestHelpWithHFlagB(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Expected crash on help flag use")
		}
	}()
	p := flaggy.NewParser("TestHelpWithHFlag")
	p.ShowHelpWithHFlag = true
	args := []string{"--help"}
	p.ParseArgs(args)
}

func TestVersionWithVFlagB(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Expected crash on version flag use")
		}
	}()
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
	p.AttachSubcommand(newSC, 1)

	// add a positional arg onto the subcommand at relative position 1
	newSC.AddPositionalValue(&positionA, "positionalA", 1, false, "This is a test positional value")

	// override os args and parse them
	os.Args = []string{"binaryName", "testSubcommand", "testPositional"}
	p.Parse()

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
	p.AttachSubcommand(newSC, 1)

	//  test what happens if you add a bad subcommand
	os.Args = []string{"test"}
	p.Parse()
}

func TestBadPositional(t *testing.T) {

	// create the argument parser
	p := flaggy.NewParser("TestBadPositional")

	// create a subcommand
	// add a positional arg into the subcommand
	var positionA string
	var err error
	p.AddPositionalValue(&positionA, "positionalA", 1, false, "This is a test positional value")

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
		p.AttachSubcommand(newSC, 1)

		// add a positional arg onto the subcommand at relative position 1
		newSC.AddPositionalValue(&positionA, "positionalA", 1, false, "This is a test positional value")

		// override os args and parse them
		os.Args = []string{"binaryName", "testSubcommand", "testPositional"}
		err := p.Parse()
		if err != nil {
			b.Fatal("Error parsing args: " + err.Error())
		}
	}

}
